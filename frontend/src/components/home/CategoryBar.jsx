import { categories } from '../../data/categories.js';
import { localized } from '../../i18n/localized.js';
import { useTranslation } from '../../i18n/useTranslation.js';
import './CategoryBar.css';

function CategoryBar() {
  const { locale } = useTranslation();

  return (
    <nav className="category-bar" aria-label="Categories">
      <ul className="category-bar__list">
        {categories.map((category) => (
          <li key={category.id}>
            <button type="button" className="category-bar__chip">
              {localized(category, 'name', locale)}
            </button>
          </li>
        ))}
      </ul>
    </nav>
  );
}

export default CategoryBar;
