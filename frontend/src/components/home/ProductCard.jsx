import { localized } from '../../i18n/localized.js';
import { useTranslation } from '../../i18n/useTranslation.js';
import './ProductCard.css';

function formatPrice(price, locale) {
  const numberLocale = locale === 'ru' ? 'ru-RU' : 'en-US';
  return new Intl.NumberFormat(numberLocale).format(price) + ' ₽';
}

function ProductCard({ product }) {
  const { locale } = useTranslation();

  return (
    <article className="product-card">
      <button type="button" className="product-card__button">
        <div
          className="product-card__image"
          style={{ backgroundColor: product.imageColor }}
          aria-hidden="true"
        />
        <div className="product-card__body">
          <h3 className="product-card__title">
            {localized(product, 'title', locale)}
          </h3>
          <div className="product-card__rating" aria-label={`Rating ${product.rating}`}>
            ★ {product.rating.toFixed(1)}
          </div>
          <div className="product-card__prices">
            <span className="product-card__price">
              {formatPrice(product.price, locale)}
            </span>
            {product.oldPrice && (
              <span className="product-card__old-price">
                {formatPrice(product.oldPrice, locale)}
              </span>
            )}
          </div>
        </div>
      </button>
    </article>
  );
}

export default ProductCard;
