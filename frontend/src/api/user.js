import { apiFetch } from './client.js';

export function getCurrentUser(token) {
  return apiFetch('/api/user/me', {
    method: 'GET',
    token,
  });
}
