window.addEventListener("load", function () {
  const productsTable = document.querySelector("#products-table");
  const compareButton = document.createElement("button");
  const productTypeSelect = document.getElementById("product-type");
  const earnFieldsDiv = document.getElementById("earn-fields");
  const swapFieldsDiv = document.getElementById("swap-fields");
  const stakeFieldsDiv = document.getElementById("stake-fields");
  const lendFieldsDiv = document.getElementById("lend-fields");

  const productList = [];

  productsTable.addEventListener("change", (event) => {
    const productCheckbox = event.target;
    const productRow = productCheckbox.closest("tr");
    const productInfo = {
      name: productRow.querySelector("td:first-child").textContent,
      purpose: productRow.querySelector("td:nth-child(2)").textContent,
      experience: productRow.querySelector("td:nth-child(3)").textContent,
      features: productRow.querySelector("td:nth-child(4)").textContent,
    };

    if (productCheckbox.checked) {
      productList.push(productInfo);
    } else {
      const index = productList.indexOf(productInfo);
      productList.splice(index, 1);
    }
  });

  compareButton.textContent = "Compare Selected";
  compareButton.style.marginLeft = "10px";
  productsTable.parentNode.insertBefore(compareButton, productsTable.nextSibling);
  compareButton.addEventListener("click", function () {
    if (productList.length < 2) {
      alert("Please select at least two products to compare");
    } else {
      let comparisonWindow = window.open("", "Comparison Window", "width=500, height=500");
      comparisonWindow.document.write("<html><head><style>");
      comparisonWindow.document.write("table {border-collapse: collapse; width: 100%; margin-top: 20px;}");
      comparisonWindow.document.write("th, td {border: 1px solid #dddddd; text-align: left; padding: 8px;}");
      comparisonWindow.document.write("th {background-color: #dddddd;}");
      comparisonWindow.document.write("h2 {text-align: center; margin-bottom: 20px;}");
      comparisonWindow.document.write("</style></head><body>");
      comparisonWindow.document.write("<h2>Product Comparison</h2>");
      comparisonWindow.document.write("<table><thead><tr><th>Product Name</th><th>Purpose</th><th>Experience Level</th><th>Features</th></tr></thead><tbody>");

      productList.forEach(function (product) {
        comparisonWindow.document.write("<tr><td>" + product.name + "</td><td>" + product.purpose + "</td><td>" + product.experience + "</td><td>" + product.features + "</td></tr>");
      });

      comparisonWindow.document.write("</tbody></table></body></html>");
    }
  });

  document.querySelectorAll(".product-link").forEach(function(element) {
    element.addEventListener("click", function(event) {
      event.preventDefault();
  
      let productWindow = window.open("", "Product Window", "width=500, height=500");
      productWindow.document.write("<html><head><style>");
      productWindow.document.write("table {border-collapse: collapse; width: 100%; margin-top: 20px;}");
      productWindow.document.write("th, td {border: 1px solid #dddddd; text-align: left; padding: 8px;}");
      productWindow.document.write("th {background-color: #dddddd;}");
      productWindow.document.write("h2 {text-align: center; margin-bottom: 20px;}");
      productWindow.document.write("</style></head><body>");
      productWindow.document.write("<h2>Product Detail</h2>");
      productWindow.document.write("<table><thead><tr><th>Product Name</th><th>Purpose</th><th>Experience Level</th><th>Features</th></tr></thead><tbody>");
      productWindow.document.write("<tr><td>" + element.text + "</td><td>" + element.parentElement.nextElementSibling.textContent + "</td><td>" + element.parentElement.nextElementSibling.nextElementSibling.textContent + "</td><td>" + element.parentElement.nextElementSibling.nextElementSibling.nextElementSibling.textContent + "</td></tr>");
      productWindow.document.write("</tbody></table></body></html>");
    });
  });

  productTypeSelect.addEventListener("change", (event) => {
    if (event.target.value === "earn") {
      earnFieldsDiv.style.display = "block";
    } else {
      earnFieldsDiv.style.display = "none";
    }

    if (event.target.value === "swap") {
      swapFieldsDiv.style.display = "block";
    } else {
      swapFieldsDiv.style.display = "none";
    }

    if (event.target.value === "stake") {
      stakeFieldsDiv.style.display = "block";
    } else {
      stakeFieldsDiv.style.display = "none";
    }

    if (event.target.value === "lend") {
      lendFieldsDiv.style.display = "block";
    } else {
      lendFieldsDiv.style.display = "none";
    }
  });

});
