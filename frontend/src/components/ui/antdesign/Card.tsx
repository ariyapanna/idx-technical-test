import React from 'react';
import { Card as AntCard, type CardProps as AntCardProps, Skeleton, Typography } from 'antd';
import { colors, spacing } from '../../../theme/design-system';

const { Title } = Typography;

interface CustomCardProps extends AntCardProps {
  glass?: boolean;
  bordered?: boolean;
  subtitle?: string;
  loading?: boolean;
  children?: React.ReactNode;
}

const Card: React.FC<CustomCardProps> = ({
  title,
  subtitle,
  children,
  glass = false,
  bordered = true,
  loading = false,
  style,
  styles,
  ...rest
}) => {
  const { bordered: _, ...antCardProps } = rest as any;

  const glassStyle: React.CSSProperties = glass ? {
    background: 'rgba(255, 255, 255, 0.7)',
    backdropFilter: 'blur(12px)',
    WebkitBackdropFilter: 'blur(12px)',
    border: '1px solid rgba(255, 255, 255, 0.3)',
    boxShadow: '0 8px 32px 0 rgba(31, 38, 135, 0.07)',
  } : {
    boxShadow: '0 4px 6px -1px rgb(0 0 0 / 0.05), 0 2px 4px -2px rgb(0 0 0 / 0.1)',
  };

  const cardTitle = title || subtitle ? (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '4px' }}>
      {title && <Title level={5} style={{ margin: 0, fontWeight: 700 }}>{title}</Title>}
      {subtitle && <Typography.Text type="secondary" style={{ fontSize: '13px' }}>{subtitle}</Typography.Text>}
    </div>
  ) : undefined;

  const mergedStyles = typeof styles === 'function' ? styles : {
    ...styles,
    body: {
      padding: spacing.lg,
      ...(styles as any)?.body,
    },
  };

  return (
    <AntCard
      title={cardTitle}
      variant={bordered ? 'outlined' : 'borderless'}
      style={{
        borderRadius: 16,
        overflow: 'hidden',
        border: (bordered && !glass) ? `1px solid ${colors.neutral.border}` : 'none',
        ...glassStyle,
        ...style,
      }}
      styles={mergedStyles as any}
      {...antCardProps}
    >

      {loading ? (
        <Skeleton active paragraph={{ rows: 3 }} />
      ) : (
        children
      )}
    </AntCard>
  );
};



export default Card;