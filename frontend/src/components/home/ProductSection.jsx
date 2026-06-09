import { useTranslation } from '../../i18n/useTranslation.js';
import ProductCard from './ProductCard.jsx';
import './ProductSection.css';

function ProductSection({ titleKey, products }) {
  const { t } = useTranslation();

  return (
    <section className="product-section">
      <h2 className="product-section__title">{t(titleKey)}</h2>
      <div className="product-section__grid">
        {products.map((product) => (
          <ProductCard key={product.id} product={product} />
        ))}
      </div>
    </section>
  );
}

export default ProductSection;
