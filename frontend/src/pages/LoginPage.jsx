import { useState } from 'react';
import { Link, Navigate, useNavigate } from 'react-router-dom';
import { login as loginRequest } from '../api/auth.js';
import { ApiError } from '../api/client.js';
import { useAuth } from '../auth/useAuth.js';
import { useTranslation } from '../i18n/useTranslation.js';
import { validateEmail, validateRequired } from './formValidation.js';
import './LoginPage.css';

function LoginPage() {
  const { t } = useTranslation();
  const { isAuthenticated, login } = useAuth();
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [fieldErrors, setFieldErrors] = useState({});
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  if (isAuthenticated) {
    return <Navigate to="/account" replace />;
  }

  const setFieldError = (field, message) => {
    setFieldErrors((prev) => ({ ...prev, [field]: message }));
  };

  const clearFieldError = (field) => {
    setFieldErrors((prev) => {
      if (!prev[field]) {
        return prev;
      }
      const next = { ...prev };
      delete next[field];
      return next;
    });
  };

  const validateForm = () => {
    const nextErrors = {
      email: validateEmail(email, t),
      password: validateRequired(password, t),
    };
    setFieldErrors(nextErrors);
    return !nextErrors.email && !nextErrors.password;
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    setError('');

    if (!validateForm()) {
      return;
    }

    setLoading(true);

    try {
      const result = await loginRequest({ email: email.trim(), password });
      const accessToken = result.data?.accessToken;
      if (!accessToken) {
        throw new ApiError(t('authErrorGeneric'));
      }
      login(accessToken);
      navigate('/account');
    } catch (err) {
      setError(err instanceof ApiError ? err.message : t('authErrorGeneric'));
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="login-page">
      <h1 className="login-page__title">{t('loginTitle')}</h1>

      <section className="login-page__card">
        <form className="login-page__form" onSubmit={handleSubmit} noValidate>
          <label className={`login-page__field${fieldErrors.email ? ' login-page__field--error' : ''}`}>
            <span className="login-page__label">{t('email')}</span>
            <input
              type="email"
              name="email"
              className={`login-page__input${fieldErrors.email ? ' login-page__input--error' : ''}`}
              placeholder={t('emailPlaceholder')}
              autoComplete="email"
              value={email}
              aria-invalid={Boolean(fieldErrors.email)}
              aria-describedby={fieldErrors.email ? 'login-email-error' : undefined}
              onChange={(event) => {
                setEmail(event.target.value);
                clearFieldError('email');
              }}
              onBlur={() => setFieldError('email', validateEmail(email, t))}
            />
            {fieldErrors.email ? (
              <span id="login-email-error" className="login-page__field-error" role="alert">
                <span className="login-page__field-error-icon" aria-hidden="true">!</span>
                {fieldErrors.email}
              </span>
            ) : null}
          </label>

          <label className={`login-page__field${fieldErrors.password ? ' login-page__field--error' : ''}`}>
            <span className="login-page__label">{t('password')}</span>
            <input
              type="password"
              name="password"
              className={`login-page__input${fieldErrors.password ? ' login-page__input--error' : ''}`}
              placeholder={t('passwordPlaceholder')}
              autoComplete="current-password"
              value={password}
              aria-invalid={Boolean(fieldErrors.password)}
              aria-describedby={fieldErrors.password ? 'login-password-error' : undefined}
              onChange={(event) => {
                setPassword(event.target.value);
                clearFieldError('password');
              }}
              onBlur={() => setFieldError('password', validateRequired(password, t))}
            />
            {fieldErrors.password ? (
              <span id="login-password-error" className="login-page__field-error" role="alert">
                <span className="login-page__field-error-icon" aria-hidden="true">!</span>
                {fieldErrors.password}
              </span>
            ) : null}
          </label>

          <a href="#" className="login-page__forgot">
            {t('forgotPassword')}
          </a>

          {error ? <p className="login-page__error" role="alert">{error}</p> : null}

          <button type="submit" className="login-page__submit" disabled={loading}>
            {loading ? t('signingIn') : t('signIn')}
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
