import { apiFetch } from './client.js';

export function getCurrentUser(token) {
  return apiFetch('/api/user/me', {
    method: 'GET',
    token,
  });
}

export function getProfileImage(token) {
  return apiFetch('/api/user/profile-image/me', {
    method: 'GET',
    token,
    responseType: 'blob',
  });
}
