import i18next from 'i18next';
import LanguageDetector from 'i18next-browser-languagedetector';

import { enNZTranslation as enNZOrderDetails } from '../components/pages/order/OrderDetails';
import { enNZTranslation as enNZOrderList } from '../components/pages/order/OrderList';
import { enNZTranslation as enNZSubmitOrder } from '../components/pages/order/SubmitOrder';

i18next.use(LanguageDetector).init({
  interpolation: {
    escapeValue: false,
  },
  lng: 'en_NZ',
  resources: {
    en_NZ: {
      translation: {
        omiga: { title: 'Omiga' },

        orderList: enNZOrderList,
        orderDetails: enNZOrderDetails,
        submitOrder: enNZSubmitOrder,
      },
    },
  },
});

export default i18next;
