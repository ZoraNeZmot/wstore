import { useEffect, useState } from 'react';
import { Link, Navigate, useNavigate } from 'react-router-dom';
import { ApiError } from '../api/client.js';
import { getFriendlyError } from '../api/friendlyError.js';
import ErrorBanner from '../components/ErrorBanner.jsx';
import { useAuth } from '../auth/useAuth.js';
import { useTranslation } from '../i18n/useTranslation.js';
import './AccountPage.css';

const accountNavLinks = [
  { key: 'accountNavProfile', href: '#' },
  { key: 'accountNavOrders', href: '#' },
  { key: 'accountNavWishlist', href: '#' },
  { key: 'accountNavAddresses', href: '#' },
  { key: 'accountNavPayments', href: '#' },
  { key: 'accountNavSecurity', href: '#' },
];

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
  const [imageUrl, setImageUrl] = useState(null);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!isAuthenticated) {
      return;
    }

    let cancelled = false;
    let objectUrl = null;

    async function loadUser() {
      setLoading(true);
      setError(null);
      setImageUrl(null);

      try {
        const result = await authFetch('/api/user/me');
        if (cancelled) {
          return;
        }
        setUser(result.data);

        try {
          const imageResult = await authFetch('/api/user/profile-image/me', {
            responseType: 'blob',
          });
          if (!cancelled) {
            objectUrl = URL.createObjectURL(imageResult.blob);
            setImageUrl(objectUrl);
          }
        } catch {
          // Profile image is optional; keep account data visible without it.
        }
      } catch (err) {
        if (!cancelled) {
          setError(getFriendlyError(err, t));
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
      if (objectUrl) {
        URL.revokeObjectURL(objectUrl);
      }
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

      <div className="account-page__layout">
        <aside className="account-page__sidebar">
          <section className="account-page__card">
            {loading ? (
              <p className="account-page__status">{t('loadingUser')}</p>
            ) : null}

            {!loading && error ? (
              <ErrorBanner error={error} />
            ) : null}

            {!loading && !error && user ? (
              <>
                <div className="account-page__avatar">
                  {imageUrl ? (
                    <img
                      src={imageUrl}
                      alt={t('accountAvatarAlt')}
                      className="account-page__avatar-image"
                    />
                  ) : (
                    <div className="account-page__avatar-fallback" aria-hidden="true">
                      {formatDisplayName(user.name).charAt(0) || '?'}
                    </div>
                  )}
                </div>

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
              </>
            ) : null}

            <nav className="account-page__nav" aria-label={t('accountNavLabel')}>
              <ul className="account-page__nav-list">
                {accountNavLinks.map((item) => (
                  <li key={item.key}>
                    <a href={item.href} className="account-page__nav-link">
                      {t(item.key)}
                    </a>
                  </li>
                ))}
              </ul>
            </nav>

            <button type="button" className="account-page__sign-out" onClick={handleSignOut}>
              {t('signOut')}
            </button>
          </section>
        </aside>
      </div>

      <Link to="/" className="account-page__back">
        {t('backToHome')}
      </Link>
    </div>
  );
}

export default AccountPage;
