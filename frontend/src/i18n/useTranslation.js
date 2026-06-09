import { useLanguage } from './useLanguage.js';

export function useTranslation() {
  const { locale, locales } = useLanguage();
  const strings = locales[locale];

  const t = (key) => strings[key] ?? key;

  return { t, locale };
}
