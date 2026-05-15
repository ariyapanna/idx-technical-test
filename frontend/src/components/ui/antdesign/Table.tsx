import React from 'react';
import { Table as AntTable } from 'antd';
import type { TableProps as AntTableProps } from 'antd';

interface CustomTableProps<T> {
  dataSource: T[];
  columns: any[];
  rowKey: string | ((record: T) => string | number);
  loading?: boolean;
  pagination?: AntTableProps<T>['pagination'];
  onChange?: AntTableProps<T>['onChange'];
  scroll?: AntTableProps<T>['scroll'];
  size?: 'small' | 'middle' | 'large';
  onRow?: (record: T) => React.HTMLAttributes<any>;
}

const Table = <T extends object>(props: CustomTableProps<T>) => {
  return <AntTable {...props} />;
};

export default Table;