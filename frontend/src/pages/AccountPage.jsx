import { useEffect, useState } from 'react';
import { Link, Navigate, useNavigate } from 'react-router-dom';
import { ApiError } from '../api/client.js';
import { useAuth } from '../auth/useAuth.js';
import { useTranslation } from '../i18n/useTranslation.js';
import './AccountPage.css';

function formatDisplayName(name) {
  if (!name) {
    return '';
  }
  const lower = name.toLocaleLowerCase();
  return lower.charAt(0).toLocaleUpperCase() + lower.slice(1);
}

function AccountPage() {
  const { t } = useTranslation();
  const { isAuthenticated, authFetch, logout } = useAuth();
  const navigate = useNavigate();
  const [user, setUser] = useState(null);
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!isAuthenticated) {
      return;
    }

    let cancelled = false;

    async function loadUser() {
      setLoading(true);
      setError('');

      try {
        const result = await authFetch('/api/user/me');
        if (!cancelled) {
          setUser(result.data);
        }
      } catch (err) {
        if (!cancelled) {
          setError(err instanceof ApiError ? err.message : t('authErrorGeneric'));
          if (err instanceof ApiError && err.status === 401) {
            logout();
          }
        }
      } finally {
        if (!cancelled) {
          setLoading(false);
        }
      }
    }

    loadUser();

    return () => {
      cancelled = true;
    };
    // t is only used for a fallback error string; omit to avoid refetch loops
  }, [isAuthenticated, authFetch, logout]);

  if (!isAuthenticated) {
    return <Navigate to="/login" replace />;
  }

  const handleSignOut = async () => {
    await logout();
    navigate('/login');
  };

  return (
    <div className="account-page">
      <h1 className="account-page__title">{t('accountTitle')}</h1>

      <section className="account-page__card">
        {loading ? (
          <p className="account-page__status">{t('loadingUser')}</p>
        ) : null}

        {!loading && error ? (
          <p className="account-page__error" role="alert">{error}</p>
        ) : null}

        {!loading && !error && user ? (
          <dl className="account-page__details">
            <div className="account-page__row">
              <dt className="account-page__label">{t('name')}</dt>
              <dd className="account-page__value">{formatDisplayName(user.name)}</dd>
            </div>
            <div className="account-page__row">
              <dt className="account-page__label">{t('email')}</dt>
              <dd className="account-page__value">{user.email}</dd>
            </div>
            <div className="account-page__row">
              <dt className="account-page__label">{t('userId')}</dt>
              <dd className="account-page__value">{user.id}</dd>
            </div>
          </dl>
        ) : null}

        <button type="button" className="account-page__sign-out" onClick={handleSignOut}>
          {t('signOut')}
        </button>
      </section>

      <Link to="/" className="account-page__back">
        {t('backToHome')}
      </Link>
    </div>
  );
}

export default AccountPage;
