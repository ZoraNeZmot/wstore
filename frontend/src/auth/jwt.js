const EXPIRY_SKEW_MS = 30_000;

function base64UrlDecode(value) {
  const padded = value.replace(/-/g, '+').replace(/_/g, '/');
  const padLength = (4 - (padded.length % 4)) % 4;
  const base64 = padded + '='.repeat(padLength);
  return atob(base64);
}

export function getTokenExpiryMs(token) {
  if (!token) {
    return null;
  }

  try {
    const payloadPart = token.split('.')[1];
    if (!payloadPart) {
      return null;
    }
    const payload = JSON.parse(base64UrlDecode(payloadPart));
    return typeof payload.exp === 'number' ? payload.exp * 1000 : null;
  } catch {
    return null;
  }
}

export function isTokenExpired(token, skewMs = EXPIRY_SKEW_MS) {
  const expiryMs = getTokenExpiryMs(token);
  if (expiryMs == null) {
    return true;
  }
  return Date.now() >= expiryMs - skewMs;
}
