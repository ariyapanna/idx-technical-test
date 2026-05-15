import type { ThemeConfig } from 'antd';

/**
 * Design System Constants
 * Use these for consistent spacing, colors, and typography across the app.
 */

export const colors = {
  primary: '#6366f1', // Modern Indigo
  primaryHover: '#4f46e5',
  primaryLight: '#eef2ff',

  success: '#10b981', // Emerald
  warning: '#f59e0b', // Amber
  error: '#ef4444',   // Rose
  info: '#3b82f6',    // Blue

  neutral: {
    title: '#1e293b',    // Slate 900
    text: '#475569',     // Slate 700
    secondary: '#64748b', // Slate 500
    muted: '#94a3b8',     // Slate 400
    border: '#e2e8f0',    // Slate 200
    bgLayout: '#f8fafc',  // Slate 50
    bgComponent: '#ffffff',
  },
};

export const spacing = {
  xs: 4,
  sm: 8,
  md: 16,
  lg: 24,
  xl: 32,
  xxl: 48,
  pagePadding: 24,
};

export const typography = {
  fontFamily: "'Plus Jakarta Sans', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif",
  fontSize: {
    xs: 12,
    sm: 14,
    md: 16,
    lg: 18,
    xl: 20,
    xxl: 24,
    heading: 32,
  },
  fontWeight: {
    regular: 400,
    medium: 500,
    semibold: 600,
    bold: 700,
  },
};

/**
 * Ant Design Theme Configuration
 * This object is passed to <ConfigProvider theme={antdTheme}>
 */
export const antdTheme: ThemeConfig = {
  token: {
    colorPrimary: colors.primary,
    colorSuccess: colors.success,
    colorWarning: colors.warning,
    colorError: colors.error,
    colorInfo: colors.info,

    colorTextBase: colors.neutral.text,
    colorTextHeading: colors.neutral.title,
    colorBgLayout: colors.neutral.bgLayout,
    colorBgContainer: colors.neutral.bgComponent,
    colorBorder: colors.neutral.border,

    borderRadius: 8,
    fontFamily: typography.fontFamily,
    fontSize: typography.fontSize.sm,

    // Spacing
    marginXS: spacing.xs,
    marginSM: spacing.sm,
    margin: spacing.md,
    marginLG: spacing.lg,
    paddingXS: spacing.xs,
    paddingSM: spacing.sm,
    padding: spacing.md,
    paddingLG: spacing.lg,
  },
  components: {
    Button: {
      borderRadius: 8,
      controlHeight: 38,
      fontWeight: 600,
      boxShadow: 'none',
    },
    Card: {
      borderRadiusLG: 16,
      boxShadowTertiary: '0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px -1px rgba(0, 0, 0, 0.1)',
    },
    Input: {
      controlHeight: 38,
      borderRadius: 8,
    },
    Select: {
      controlHeight: 38,
      borderRadius: 8,
    },
    Table: {
      borderRadius: 12,
    },
    Layout: {
      colorBgHeader: '#ffffff',
      colorBgBody: colors.neutral.bgLayout,
    },
    Menu: {
      itemBorderRadius: 8,
      itemMarginInline: 8,
    }
  },
};
