import { useState, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import Layout from './components/Layout';
import Sidebar from './components/Sidebar';
import { GeneralSettingsPage, AdvancedSettingsPage } from './pages/SettingsPages';
import AboutPage from './pages/AboutPage';
import { GetConfig } from '../wailsjs/go/main/App';

function App() {
  const { i18n } = useTranslation();
  const [currentPage, setCurrentPage] = useState('settings-general');

  useEffect(() => {
    const initLanguage = async () => {
      try {
        const config = await GetConfig();
        if (config && config.application && config.application.language) {
          i18n.changeLanguage(config.application.language);
        }
      } catch (error) {
        console.error('Failed to load language setting:', error);
      }
    };
    
    initLanguage();
  }, [i18n]);

  const renderPage = () => {
    switch (currentPage) {
      case 'settings-general':
        return <GeneralSettingsPage />;
      case 'settings-advanced':
        return <AdvancedSettingsPage />;
      case 'about':
        return <AboutPage />;
      default:
        return <GeneralSettingsPage />;
    }
  };

  return (
    <Layout sidebar={<Sidebar onNavigate={setCurrentPage} />}>
      {renderPage()}
    </Layout>
  );
}

export default App;
