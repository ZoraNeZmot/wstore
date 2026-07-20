import { ApiError } from './client.js';

function textOf(value) {
  if (value == null) {
    return '';
  }
  if (typeof value === 'string') {
    return value;
  }
  try {
    return JSON.stringify(value);
  } catch {
    return String(value);
  }
}

function includesAny(haystack, needles) {
  const text = haystack.toLowerCase();
  return needles.some((needle) => text.includes(needle.toLowerCase()));
}

function isNetworkError(err) {
  if (!err || err instanceof ApiError) {
    return false;
  }
  const message = textOf(err.message).toLowerCase();
  return (
    err.name === 'TypeError'
    || includesAny(message, ['failed to fetch', 'networkerror', 'load failed', 'network request failed'])
  );
}

/**
 * Maps API / network errors to user-facing title + recommendation.
 * @returns {{ title: string, tip: string }}
 */
export function getFriendlyError(err, t) {
  if (isNetworkError(err)) {
    return {
      title: t('errNetworkTitle'),
      tip: t('errNetworkTip'),
    };
  }

  if (!(err instanceof ApiError)) {
    return {
      title: t('errGenericTitle'),
      tip: t('errGenericTip'),
    };
  }

  const status = err.status;
  const message = textOf(err.message);
  const details = textOf(err.details);
  const blob = `${message} ${details}`;

  if (status === 409 || includesAny(blob, ['already exists', 'user already exists', 'conflict'])) {
    return {
      title: t('errUserExistsTitle'),
      tip: t('errUserExistsTip'),
    };
  }

  if (
    status === 401
    && includesAny(blob, ['invalid email or password', 'user not found'])
  ) {
    return {
      title: t('errInvalidCredentialsTitle'),
      tip: t('errInvalidCredentialsTip'),
    };
  }

  if (
    status === 401
    || includesAny(blob, ['unauthorized', 'invalid token', 'invalid refresh', 'failed to refresh'])
  ) {
    return {
      title: t('errSessionTitle'),
      tip: t('errSessionTip'),
    };
  }

  if (status === 400 || includesAny(blob, ['bad request', 'failed to parse', 'failed to read'])) {
    return {
      title: t('errBadRequestTitle'),
      tip: t('errBadRequestTip'),
    };
  }

  if (status === 403) {
    return {
      title: t('errForbiddenTitle'),
      tip: t('errForbiddenTip'),
    };
  }

  if (status === 404) {
    return {
      title: t('errNotFoundTitle'),
      tip: t('errNotFoundTip'),
    };
  }

  if (status === 429) {
    return {
      title: t('errTooManyTitle'),
      tip: t('errTooManyTip'),
    };
  }

  if (status === 502 || status === 503 || status === 504) {
    return {
      title: t('errUnavailableTitle'),
      tip: t('errUnavailableTip'),
    };
  }

  if (status >= 500 || includesAny(blob, ['server error', 'cannot '])) {
    return {
      title: t('errServerTitle'),
      tip: t('errServerTip'),
    };
  }

  return {
    title: t('errGenericTitle'),
    tip: t('errGenericTip'),
  };
}
