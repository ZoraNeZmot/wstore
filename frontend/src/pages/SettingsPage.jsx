import { Link } from 'react-router-dom';
import { useLanguage } from '../i18n/useLanguage.js';
import { useTranslation } from '../i18n/useTranslation.js';
import './SettingsPage.css';

const languageOptions = [
  { value: 'en', labelKey: 'languageEn' },
  { value: 'ru', labelKey: 'languageRu' },
];

function SettingsPage() {
  const { locale, setLocale } = useLanguage();
  const { t } = useTranslation();

  return (
    <div className="settings-page">
      <h1 className="settings-page__title">{t('settingsTitle')}</h1>

      <section className="settings-page__section">
        <h2 className="settings-page__section-title">{t('language')}</h2>
        <div className="settings-page__options" role="radiogroup" aria-label={t('language')}>
          {languageOptions.map((option) => (
            <label key={option.value} className="settings-page__option">
              <input
                type="radio"
                name="language"
                value={option.value}
                checked={locale === option.value}
                onChange={() => setLocale(option.value)}
              />
              <span>{t(option.labelKey)}</span>
            </label>
          ))}
        </div>
      </section>

      <Link to="/" className="settings-page__back">
        {t('backToHome')}
      </Link>
    </div>
  );
}

export default SettingsPage;
