import './ErrorBanner.css';

function ErrorBanner({ error }) {
  if (!error?.title) {
    return null;
  }

  return (
    <div className="error-banner" role="alert">
      <div className="error-banner__icon" aria-hidden="true">!</div>
      <div className="error-banner__body">
        <p className="error-banner__title">{error.title}</p>
        {error.tip ? <p className="error-banner__tip">{error.tip}</p> : null}
      </div>
    </div>
  );
}

export default ErrorBanner;
