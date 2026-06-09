export function localized(item, field, locale) {
  const suffix = locale === 'ru' ? 'Ru' : 'En';
  return item[`${field}${suffix}`];
}
