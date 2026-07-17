const EMAIL_PATTERN = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

export function validateRequired(value, t) {
  if (!String(value ?? '').trim()) {
    return t('fieldRequired');
  }
  return '';
}

export function validateEmail(value, t) {
  const trimmed = String(value ?? '').trim();
  if (!trimmed) {
    return t('fieldRequired');
  }
  if (!EMAIL_PATTERN.test(trimmed)) {
    return t('emailInvalid');
  }
  return '';
}

export function validatePasswordMatch(password, confirmPassword, t) {
  if (!String(confirmPassword ?? '').trim()) {
    return t('fieldRequired');
  }
  if (password !== confirmPassword) {
    return t('passwordMismatch');
  }
  return '';
}
