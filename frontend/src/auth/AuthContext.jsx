import { useCallback, useMemo, useRef, useState } from 'react';
import { logout as logoutRequest, refresh as refreshTokenRequest } from '../api/auth.js';
import { apiFetch, ApiError } from '../api/client.js';
import { AuthContext } from './authContext.js';
import { isTokenExpired } from './jwt.js';

const STORAGE_KEY = 'wstore-access-token';

function getStoredToken() {
  return localStorage.getItem(STORAGE_KEY);
}

export function AuthProvider({ children }) {
  const [accessToken, setAccessTokenState] = useState(getStoredToken);
  const tokenRef = useRef(accessToken);
  const refreshPromiseRef = useRef(null);

  const setAccessToken = useCallback((token) => {
    tokenRef.current = token;
    setAccessTokenState(token);
    if (token) {
      localStorage.setItem(STORAGE_KEY, token);
    } else {
      localStorage.removeItem(STORAGE_KEY);
    }
  }, []);

  const login = useCallback((token) => {
    setAccessToken(token);
  }, [setAccessToken]);

  const logout = useCallback(async () => {
    try {
      await logoutRequest();
    } catch {
      // Clear local session even if the API call fails.
    } finally {
      setAccessToken(null);
    }
  }, [setAccessToken]);

  const refreshAccessToken = useCallback(async () => {
    if (!refreshPromiseRef.current) {
      refreshPromiseRef.current = refreshTokenRequest()
        .then((result) => {
          const nextToken = result.data?.accessToken;
          if (!nextToken) {
            throw new ApiError('Failed to refresh session', { status: 401 });
          }
          setAccessToken(nextToken);
          return nextToken;
        })
        .catch((error) => {
          setAccessToken(null);
          throw error;
        })
        .finally(() => {
          refreshPromiseRef.current = null;
        });
    }
    return refreshPromiseRef.current;
  }, [setAccessToken]);

  const ensureAccessToken = useCallback(async () => {
    const currentToken = tokenRef.current;
    if (!currentToken) {
      throw new ApiError('Unauthorized', { status: 401 });
    }
    if (!isTokenExpired(currentToken)) {
      return currentToken;
    }
    return refreshAccessToken();
  }, [refreshAccessToken]);

  const authFetch = useCallback(async (path, options = {}) => {
    const { skipAuth = false, ...rest } = options;

    if (skipAuth) {
      return apiFetch(path, {
        ...rest,
        skipAuth: true,
      });
    }

    const token = await ensureAccessToken();
    return apiFetch(path, {
      ...rest,
      skipAuth: false,
      token,
    });
  }, [ensureAccessToken]);

  const value = useMemo(() => ({
    accessToken,
    isAuthenticated: Boolean(accessToken),
    login,
    logout,
    authFetch,
  }), [accessToken, login, logout, authFetch]);

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
}
