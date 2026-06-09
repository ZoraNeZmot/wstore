import { Link } from 'react-router-dom';
import { useTranslation } from '../i18n/useTranslation.js';
import './LoginPage.css';

function LoginPage() {
  const { t } = useTranslation();

  const handleSubmit = (event) => {
    event.preventDefault();
  };

  return (
    <div className="login-page">
      <h1 className="login-page__title">{t('loginTitle')}</h1>

      <section className="login-page__card">
        <form className="login-page__form" onSubmit={handleSubmit}>
          <label className="login-page__field">
            <span className="login-page__label">{t('email')}</span>
            <input
              type="email"
              name="email"
              className="login-page__input"
              placeholder={t('emailPlaceholder')}
              autoComplete="email"
            />
          </label>

          <label className="login-page__field">
            <span className="login-page__label">{t('password')}</span>
            <input
              type="password"
              name="password"
              className="login-page__input"
              placeholder={t('passwordPlaceholder')}
              autoComplete="current-password"
            />
          </label>

          <a href="#" className="login-page__forgot">
            {t('forgotPassword')}
          </a>

          <button type="submit" className="login-page__submit">
            {t('signIn')}
          </button>
        </form>

        <div className="login-page__divider">
          <span>{t('or')}</span>
        </div>

        <button type="button" className="login-page__google">
          {t('signInWithGoogle')}
        </button>
      </section>

      <p className="login-page__register">
        {t('noAccount')}{' '}
        <Link to="/register" className="login-page__register-link">
          {t('register')}
        </Link>
      </p>

      <Link to="/" className="login-page__back">
        {t('backToHome')}
      </Link>
    </div>
  );
}

export default LoginPage;
