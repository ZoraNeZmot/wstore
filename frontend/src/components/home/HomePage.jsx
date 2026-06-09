import { popularProducts, dealProducts } from '../../data/products.js';
import HeroBanner from './HeroBanner.jsx';
import CategoryBar from './CategoryBar.jsx';
import ProductSection from './ProductSection.jsx';

function HomePage() {
  return (
    <>
      <HeroBanner />
      <CategoryBar />
      <ProductSection titleKey="popular" products={popularProducts} />
      <ProductSection titleKey="deals" products={dealProducts} />
    </>
  );
}

export default HomePage;
