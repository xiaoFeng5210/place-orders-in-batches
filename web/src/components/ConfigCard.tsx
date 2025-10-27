import React from 'react';
import { Card, Switch, Button, Tag, Space, Popconfirm } from 'antd';
import { DeleteOutlined, EditOutlined, ShopOutlined } from '@ant-design/icons';
import type { OrderConfig } from '../types';

interface ConfigCardProps {
  config: OrderConfig;
  onToggleMock: (dealerId: string, enabled: boolean) => void;
  onDelete: (dealerId: string) => void;
  onEdit?: (config: OrderConfig) => void;
}

const ConfigCard: React.FC<ConfigCardProps> = ({
  config,
  onToggleMock,
  onDelete,
  onEdit
}) => {
  const handleMockToggle = (checked: boolean) => {
    onToggleMock(config.dealer, checked);
  };

  const handleDelete = () => {
    onDelete(config.dealer);
  };

  return (
    <Card
      className="w-full shadow-sm hover:shadow-md transition-shadow duration-200"
      title={
        <div className="flex items-center gap-2">
          <ShopOutlined className="text-blue-500" />
          <span className="font-medium">{config.dealer_name || '未命名门店'}</span>
          <Tag color={config.enable_mock ? 'green' : 'default'}>
            {config.enable_mock ? '已启用' : '已停用'}
          </Tag>
        </div>
      }
      extra={
        <Space>
          <Switch
            checked={config.enable_mock}
            onChange={handleMockToggle}
            checkedChildren="启用"
            unCheckedChildren="停用"
          />
          {onEdit && (
            <Button
              type="text"
              icon={<EditOutlined />}
              onClick={() => onEdit(config)}
              className="text-gray-500 hover:text-blue-500"
            />
          )}
          <Popconfirm
            title="确认删除"
            description="确定要删除这个门店配置吗？"
            onConfirm={handleDelete}
            okText="确认"
            cancelText="取消"
          >
            <Button
              type="text"
              danger
              icon={<DeleteOutlined />}
              className="text-gray-500 hover:text-red-500"
            />
          </Popconfirm>
        </Space>
      }
    >
      <div className="space-y-3">
        <div className="flex items-center justify-between">
          <span className="text-gray-600">门店ID:</span>
          <span className="font-mono text-sm bg-gray-100 px-2 py-1 rounded">
            {config.dealer}
          </span>
        </div>
        
        <div className="flex items-center justify-between">
          <span className="text-gray-600">订单金额:</span>
          <span className="font-semibold text-green-600">
            ¥{config.amount.toFixed(2)}
          </span>
        </div>
        
        <div className="flex items-center justify-between">
          <span className="text-gray-600">支付方式:</span>
          <span>{config.pay_type || '未设置'}</span>
        </div>
        
        <div>
          <div className="text-gray-600 mb-2">商品列表:</div>
          <div className="space-y-1">
            {config.products.map((product, index) => (
              <div
                key={index}
                className="flex items-center justify-between bg-gray-50 p-2 rounded"
              >
                <div className="flex-1">
                  <div className="font-medium">{product.name}</div>
                  <div className="text-xs text-gray-500 font-mono">
                    ID: {product.product_id}
                  </div>
                </div>
                <div className="text-right">
                  <span className="bg-blue-100 text-blue-800 px-2 py-1 rounded text-sm">
                    x{product.quantity}
                  </span>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </Card>
  );
};

export default ConfigCard;
