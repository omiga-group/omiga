import dataclasses
import os
import re
import argparse
from typing import Dict, List, Callable
from enum import Enum

regex = re.compile(r"\$(VENUE)\@(PAS|LOW|UP)\$", re.IGNORECASE)


class Case(Enum):
    PascalCase = "PAS"
    LowerCase = "LOW"
    UpperCase = "UP"
    Unknown = "WTF"


def str_to_case(s: str) -> Case:
    match s.upper():
        case "PAS":
            return Case.PascalCase
        case "LOW":
            return Case.LowerCase
        case "UP":
            return Case.UpperCase
    return Case.Unknown


@dataclasses.dataclass
class Venue:
    name: str
    case: Case


def venue_to_string(venue: Venue) -> str:
    match venue.case:
        case Case.PascalCase:
            return venue.name.lower().capitalize()

        case Case.LowerCase:
            return venue.name.lower()

        case Case.UpperCase:
            return venue.name.upper()


def transformer(venue_name: str) -> Callable[[re.Match], str]:
    def replacer(matches: re.Match):
        venue = Venue(name=venue_name, case=str_to_case(matches.group(2)))
        return venue_to_string(venue)

    return replacer


def transform_text(text: str, venue_name: str) -> str:
    return re.sub(regex, transformer(venue_name=venue_name), text)


def transform_path(path: str, venue_name: str, out_dir: str) -> str:
    new_path = transform_text(text=path, venue_name=venue_name)
    return new_path.replace("templates", out_dir, 1)


def get_tree(root: str) -> List[str]:
    tree = []
    files = os.listdir(root)

    for file in files:
        if file.startswith("."):
            continue

        file_path = os.path.join(root, file)
        if os.path.isfile(file_path):
            tree.append(file_path)
        else:
            sub_tree = get_tree(file_path)
            tree = tree + sub_tree

    return tree


def transform_tree(tree: List[str], venue_name: str, out_dir: str) -> Dict[str, str]:
    tree_map = {}
    for path in tree:
        tree_map[path] = transform_path(path=path, venue_name=venue_name, out_dir=out_dir)

    return tree_map


def create_tree(tree: Dict[str, str], venue_name: str):
    for source in tree:
        target = tree[source]
        base_dir = os.path.dirname(target)
        os.makedirs(base_dir, exist_ok=True)
        with open(target, "w") as tf:
            with open(source, "r") as sf:
                text = sf.read()
                tf.write(transform_text(text=text, venue_name=venue_name))


def input_args():
    args_def = argparse.ArgumentParser()
    args_def.add_argument("-n", "--name", required=True, help="name of the venue to be generated")
    args_def.add_argument("-o", "--out", required=True, help="output directory")
    args = vars(args_def.parse_args())
    return args


def main():
    args = input_args()
    venue_name = args["name"]
    out_dir = args["out"]
    src_tree = get_tree("templates")
    tree_map = transform_tree(tree=src_tree, venue_name=venue_name, out_dir=out_dir)
    create_tree(tree=tree_map, venue_name=venue_name)


main()
