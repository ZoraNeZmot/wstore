import { apiFetch } from './client.js';

export function signup({ name, email, password }) {
  return apiFetch('/api/auth/email/signup', {
    method: 'POST',
    body: { name, email, password },
    skipAuth: true,
  });
}

export function login({ email, password }) {
  return apiFetch('/api/auth/email/login', {
    method: 'POST',
    body: { email, password },
    skipAuth: true,
  });
}

export function refresh() {
  return apiFetch('/api/auth/refresh', {
    method: 'GET',
    skipAuth: true,
  });
}

export function logout() {
  return apiFetch('/api/auth/logout', {
    method: 'POST',
    skipAuth: true,
  });
}
