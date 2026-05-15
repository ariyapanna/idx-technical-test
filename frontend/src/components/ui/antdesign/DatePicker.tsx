import React from 'react';
import { DatePicker as AntDatePicker } from 'antd';
import type { Dayjs } from 'dayjs';

interface DatePickerProps {
  value?: Dayjs | null;
  onChange?: (date: Dayjs | null, dateString: any) => void;
  style?: React.CSSProperties;
  placeholder?: string;
}


const DatePicker: React.FC<DatePickerProps> = (props) => {
  return <AntDatePicker {...props} />;
};

export default DatePicker;
