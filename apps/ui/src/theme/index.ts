import type { ThemeConfig } from 'antd';

const theme: ThemeConfig = {
  cssVar: true,
  token: {
    // Base colors
    colorPrimary: '#219EBC',
    colorInfo: '#219EBC',
    colorSuccess: '#52C41A',
    colorWarning: '#FFB703',
    colorError: '#FF4D4F',

    // Neutral colors
    colorText: '#4B5C6B',
    colorTextSecondary: '#8696AA',
    colorTextTertiary: '#BFC9D9',
    colorTextQuaternary: '#E4E8ED',

    // Background colors
    colorBgContainer: '#FFFFFF',
    colorBgElevated: '#FFFFFF',
    colorBgLayout: '#F5F7FA',
    colorBgSpotlight: '#F0F5FF',

    // Border colors
    colorBorder: '#E4E8ED',
    colorBorderSecondary: '#F0F2F5',

    // Control colors
    colorFillQuaternary: '#F5F7FA',
    colorFillTertiary: '#F0F2F5',
    colorFillSecondary: '#E8EDF2',
    colorFill: '#E4E8ED',
  },
  components: {
    Layout: {
      headerBg: '#023047',
      siderBg: '#023047',
      bodyBg: '#F5F7FA',
      triggerBg: '#154964',
    },
    Menu: {
      darkItemBg: '#023047',
      darkItemColor: 'rgba(255, 255, 255, 0.65)',
      darkItemSelectedBg: '#154964',
      darkItemSelectedColor: '#FFFFFF',
      darkItemHoverBg: 'rgba(255, 255, 255, 0.05)',
      darkItemHoverColor: '#FFFFFF',
    },
    Button: {
      colorPrimaryHover: '#1A89A5',
      colorPrimaryActive: '#17778F',
    },
  },
};

export default theme;
