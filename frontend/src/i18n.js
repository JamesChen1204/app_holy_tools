import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';

import zhTW from './locales/zh-TW/translation.json';
import zhCN from './locales/zh-CN/translation.json';
import en from './locales/en/translation.json';

i18n
  .use(initReactI18next)
  .init({
    resources: {
      'zh-TW': {
        translation: zhTW
      },
      'zh-CN': {
        translation: zhCN
      },
      'en': {
        translation: en
      }
    },
    lng: 'zh-TW',
    fallbackLng: 'zh-TW',
    debug: false,
    interpolation: {
      escapeValue: false
    },
    ns: ['translation'],
    defaultNS: 'translation'
  });

export default i18n;
