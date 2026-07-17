import { useCallback } from 'react';
import { useLanguage } from './useLanguage.js';

export function useTranslation() {
  const { locale, locales } = useLanguage();
  const strings = locales[locale];

  const t = useCallback((key) => strings[key] ?? key, [strings]);

  return { t, locale };
}
