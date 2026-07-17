import { useState } from 'react';
import { Link, Navigate, useNavigate } from 'react-router-dom';
import { signup as signupRequest } from '../api/auth.js';
import { ApiError } from '../api/client.js';
import { useAuth } from '../auth/useAuth.js';
import { useTranslation } from '../i18n/useTranslation.js';
import {
  validateEmail,
  validatePasswordMatch,
  validateRequired,
} from './formValidation.js';
import './LoginPage.css';

function RegisterPage() {
  const { t } = useTranslation();
  const { isAuthenticated } = useAuth();
  const navigate = useNavigate();
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
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
      name: validateRequired(name, t),
      email: validateEmail(email, t),
      password: validateRequired(password, t),
      confirmPassword: validatePasswordMatch(password, confirmPassword, t),
    };
    setFieldErrors(nextErrors);
    return !Object.values(nextErrors).some(Boolean);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    setError('');

    if (!validateForm()) {
      return;
    }

    setLoading(true);

    try {
      await signupRequest({ name: name.trim(), email: email.trim(), password });
      navigate('/login');
    } catch (err) {
      setError(err instanceof ApiError ? err.message : t('authErrorGeneric'));
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="login-page">
      <h1 className="login-page__title">{t('registerTitle')}</h1>

      <section className="login-page__card">
        <form className="login-page__form" onSubmit={handleSubmit} noValidate>
          <label className={`login-page__field${fieldErrors.name ? ' login-page__field--error' : ''}`}>
            <span className="login-page__label">{t('name')}</span>
            <input
              type="text"
              name="name"
              className={`login-page__input${fieldErrors.name ? ' login-page__input--error' : ''}`}
              placeholder={t('namePlaceholder')}
              autoComplete="name"
              value={name}
              aria-invalid={Boolean(fieldErrors.name)}
              aria-describedby={fieldErrors.name ? 'register-name-error' : undefined}
              onChange={(event) => {
                setName(event.target.value);
                clearFieldError('name');
              }}
              onBlur={() => setFieldError('name', validateRequired(name, t))}
            />
            {fieldErrors.name ? (
              <span id="register-name-error" className="login-page__field-error" role="alert">
                <span className="login-page__field-error-icon" aria-hidden="true">!</span>
                {fieldErrors.name}
              </span>
            ) : null}
          </label>

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
              aria-describedby={fieldErrors.email ? 'register-email-error' : undefined}
              onChange={(event) => {
                setEmail(event.target.value);
                clearFieldError('email');
              }}
              onBlur={() => setFieldError('email', validateEmail(email, t))}
            />
            {fieldErrors.email ? (
              <span id="register-email-error" className="login-page__field-error" role="alert">
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
              autoComplete="new-password"
              value={password}
              aria-invalid={Boolean(fieldErrors.password)}
              aria-describedby={fieldErrors.password ? 'register-password-error' : undefined}
              onChange={(event) => {
                setPassword(event.target.value);
                clearFieldError('password');
                if (confirmPassword) {
                  setFieldError(
                    'confirmPassword',
                    validatePasswordMatch(event.target.value, confirmPassword, t),
                  );
                }
              }}
              onBlur={() => setFieldError('password', validateRequired(password, t))}
            />
            {fieldErrors.password ? (
              <span id="register-password-error" className="login-page__field-error" role="alert">
                <span className="login-page__field-error-icon" aria-hidden="true">!</span>
                {fieldErrors.password}
              </span>
            ) : null}
          </label>

          <label className={`login-page__field${fieldErrors.confirmPassword ? ' login-page__field--error' : ''}`}>
            <span className="login-page__label">{t('confirmPassword')}</span>
            <input
              type="password"
              name="confirmPassword"
              className={`login-page__input${fieldErrors.confirmPassword ? ' login-page__input--error' : ''}`}
              placeholder={t('confirmPasswordPlaceholder')}
              autoComplete="new-password"
              value={confirmPassword}
              aria-invalid={Boolean(fieldErrors.confirmPassword)}
              aria-describedby={fieldErrors.confirmPassword ? 'register-confirm-error' : undefined}
              onChange={(event) => {
                setConfirmPassword(event.target.value);
                clearFieldError('confirmPassword');
              }}
              onBlur={() => setFieldError(
                'confirmPassword',
                validatePasswordMatch(password, confirmPassword, t),
              )}
            />
            {fieldErrors.confirmPassword ? (
              <span id="register-confirm-error" className="login-page__field-error" role="alert">
                <span className="login-page__field-error-icon" aria-hidden="true">!</span>
                {fieldErrors.confirmPassword}
              </span>
            ) : null}
          </label>

          {error ? <p className="login-page__error" role="alert">{error}</p> : null}

          <button type="submit" className="login-page__submit" disabled={loading}>
            {loading ? t('creatingAccount') : t('createAccount')}
          </button>
        </form>

        <div className="login-page__divider">
          <span>{t('or')}</span>
        </div>

        <button type="button" className="login-page__google">
          {t('signUpWithGoogle')}
        </button>
      </section>

      <p className="login-page__register">
        {t('hasAccount')}{' '}
        <Link to="/login" className="login-page__register-link">
          {t('signIn')}
        </Link>
      </p>

      <Link to="/" className="login-page__back">
        {t('backToHome')}
      </Link>
    </div>
  );
}

export default RegisterPage;
