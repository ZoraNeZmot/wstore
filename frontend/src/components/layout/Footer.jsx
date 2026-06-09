import { useTranslation } from '../../i18n/useTranslation.js';
import './Footer.css';

const footerLinkKeys = ['about', 'help', 'contacts'];

function Footer() {
  const { t } = useTranslation();

  return (
    <footer className="footer">
      <div className="footer__inner">
        <nav className="footer__links" aria-label="Footer navigation">
          {footerLinkKeys.map((key) => (
            <a key={key} href="#" className="footer__link">
              {t(key)}
            </a>
          ))}
        </nav>
        <p className="footer__copyright">
          &copy; {new Date().getFullYear()} WStore
        </p>
      </div>
    </footer>
  );
}

export default Footer;
