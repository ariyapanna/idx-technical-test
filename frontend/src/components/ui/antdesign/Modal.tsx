import React from 'react';
import { Modal as AntModal } from 'antd';

interface ModalProps {
  title?: React.ReactNode;
  open?: boolean;
  onOk?: () => void;
  onCancel?: () => void;
  okText?: string;
  cancelText?: string;
  children?: React.ReactNode;
  destroyOnHidden?: boolean;
  loading?: boolean;
  footer?: React.ReactNode;
}

const Modal: React.FC<ModalProps> = (props) => {
  return <AntModal {...props} />;
};


export default Modal;
