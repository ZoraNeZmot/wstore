import { Link } from 'react-router-dom';
import { useTranslation } from '../../i18n/useTranslation.js';
import './Header.css';

function Header() {
  const { t } = useTranslation();

  return (
    <header className="header">
      <div className="header__inner">
        <Link to="/" className="header__logo">
          WStore
        </Link>

        <div className="header__search">
          <input
            type="search"
            className="header__search-input"
            placeholder={t('search')}
            aria-label={t('search')}
          />
        </div>

        <nav className="header__nav" aria-label="Main navigation">
          <a href="#" className="header__nav-link">
            {t('catalog')}
          </a>
          <a href="#" className="header__nav-link">
            {t('cart')}
          </a>
          <Link to="/login" className="header__nav-link">
            {t('account')}
          </Link>
          <Link to="/settings" className="header__nav-link">
            {t('settings')}
          </Link>
        </nav>
      </div>
    </header>
  );
}

export default Header;
