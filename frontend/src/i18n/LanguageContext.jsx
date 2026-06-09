import { useEffect, useState } from 'react';
import { locales } from './locales/index.js';
import { LanguageContext } from './languageContext.js';

const STORAGE_KEY = 'wstore-locale';
const DEFAULT_LOCALE = 'en';

function getInitialLocale() {
  const saved = localStorage.getItem(STORAGE_KEY);
  if (saved === 'en' || saved === 'ru') {
    return saved;
  }
  return DEFAULT_LOCALE;
}

export function LanguageProvider({ children }) {
  const [locale, setLocaleState] = useState(getInitialLocale);

  const setLocale = (next) => {
    if (next !== 'en' && next !== 'ru') {
      return;
    }
    setLocaleState(next);
    localStorage.setItem(STORAGE_KEY, next);
  };

  useEffect(() => {
    document.documentElement.lang = locale;
  }, [locale]);

  const value = { locale, setLocale, locales };

  return (
    <LanguageContext.Provider value={value}>
      {children}
    </LanguageContext.Provider>
  );
}
