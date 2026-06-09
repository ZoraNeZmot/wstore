import { useTranslation } from '../../i18n/useTranslation.js';
import './HeroBanner.css';

function HeroBanner() {
  const { t } = useTranslation();

  return (
    <section className="hero-banner" aria-label="Promotional banner">
      <div className="hero-banner__content">
        <h2 className="hero-banner__title">{t('heroTitle')}</h2>
        <p className="hero-banner__subtitle">{t('heroSubtitle')}</p>
      </div>
    </section>
  );
}

export default HeroBanner;
