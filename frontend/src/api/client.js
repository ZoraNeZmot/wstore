export class ApiError extends Error {
  constructor(message, { status, code, details } = {}) {
    super(message);
    this.name = 'ApiError';
    this.status = status;
    this.code = code;
    this.details = details;
  }
}

export async function apiFetch(path, options = {}) {
  const {
    method = 'GET',
    body,
    token,
    headers: extraHeaders = {},
    skipAuth = false,
    responseType = 'json',
  } = options;

  const headers = { ...extraHeaders };

  if (body !== undefined) {
    headers['Content-Type'] = 'application/json';
  }

  if (!skipAuth && token) {
    headers.Authorization = `Bearer ${token}`;
  }

  const response = await fetch(path, {
    method,
    headers,
    credentials: 'include',
    body: body !== undefined ? JSON.stringify(body) : undefined,
  });

  const contentType = response.headers.get('content-type') || '';

  if (!response.ok) {
    let errorPayload = null;
    if (contentType.includes('application/json')) {
      errorPayload = await response.json();
    }
    const error = errorPayload?.error;
    throw new ApiError(error?.message || `Request failed (${response.status})`, {
      status: response.status,
      code: error?.code,
      details: error?.details,
    });
  }

  if (responseType === 'blob') {
    return {
      status: response.status,
      blob: await response.blob(),
    };
  }

  let payload = null;
  if (contentType.includes('application/json')) {
    payload = await response.json();
  }

  if (payload?.success === false) {
    const error = payload?.error;
    throw new ApiError(error?.message || `Request failed (${response.status})`, {
      status: response.status,
      code: error?.code,
      details: error?.details,
    });
  }

  return {
    status: response.status,
    data: payload?.data,
    raw: payload,
  };
}
