import React, { useState, useEffect } from 'react';
import {
  Modal,
  Form,
  Input,
  InputNumber,
  Button,
  Space,
  Card,
  Switch,
  message
} from 'antd';
import { PlusOutlined, DeleteOutlined } from '@ant-design/icons';
import { OrderConfig, Product } from '../types';

interface AddConfigModalProps {
  visible: boolean;
  onCancel: () => void;
  onSubmit: (config: OrderConfig) => Promise<void>;
  editingConfig?: OrderConfig | null;
}

const AddConfigModal: React.FC<AddConfigModalProps> = ({
  visible,
  onCancel,
  onSubmit,
  editingConfig
}) => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);
  const [products, setProducts] = useState<Product[]>([
    { product_id: '', name: '', quantity: 1 }
  ]);

  useEffect(() => {
    if (visible) {
      if (editingConfig) {
        // 编辑模式
        form.setFieldsValue({
          dealer: editingConfig.dealer,
          dealer_name: editingConfig.dealer_name,
          enable_mock: editingConfig.enable_mock,
          amount: editingConfig.amount,
          pay_type: editingConfig.pay_type,
        });
        setProducts(editingConfig.products.length > 0 ? editingConfig.products : [
          { product_id: '', name: '', quantity: 1 }
        ]);
      } else {
        // 新增模式
        form.resetFields();
        setProducts([{ product_id: '', name: '', quantity: 1 }]);
      }
    }
  }, [visible, editingConfig, form]);

  const handleSubmit = async () => {
    try {
      setLoading(true);
      const values = await form.validateFields();
      
      // 验证商品列表
      const validProducts = products.filter(p => 
        p.product_id.trim() && p.name.trim() && p.quantity > 0
      );
      
      if (validProducts.length === 0) {
        message.error('请至少添加一个有效的商品');
        return;
      }

      const config: OrderConfig = {
        dealer: values.dealer,
        dealer_name: values.dealer_name,
        enable_mock: values.enable_mock || false,
        products: validProducts,
        amount: values.amount || 0,
        pay_type: values.pay_type || '',
      };

      await onSubmit(config);
      handleCancel();
      message.success(editingConfig ? '更新配置成功' : '添加配置成功');
    } catch (error) {
      console.error('提交失败:', error);
      message.error('操作失败，请重试');
    } finally {
      setLoading(false);
    }
  };

  const handleCancel = () => {
    form.resetFields();
    setProducts([{ product_id: '', name: '', quantity: 1 }]);
    onCancel();
  };

  const addProduct = () => {
    setProducts([...products, { product_id: '', name: '', quantity: 1 }]);
  };

  const removeProduct = (index: number) => {
    if (products.length > 1) {
      setProducts(products.filter((_, i) => i !== index));
    }
  };

  const updateProduct = (index: number, field: keyof Product, value: string | number) => {
    const newProducts = [...products];
    newProducts[index] = { ...newProducts[index], [field]: value };
    setProducts(newProducts);
  };

  return (
    <Modal
      title={editingConfig ? '编辑门店配置' : '添加门店配置'}
      open={visible}
      onCancel={handleCancel}
      width={800}
      footer={[
        <Button key="cancel" onClick={handleCancel}>
          取消
        </Button>,
        <Button
          key="submit"
          type="primary"
          loading={loading}
          onClick={handleSubmit}
        >
          {editingConfig ? '更新' : '添加'}
        </Button>,
      ]}
    >
      <Form
        form={form}
        layout="vertical"
        className="mt-4"
      >
        <div className="grid grid-cols-2 gap-4">
          <Form.Item
            label="门店ID"
            name="dealer"
            rules={[{ required: true, message: '请输入门店ID' }]}
          >
            <Input placeholder="请输入门店ID" />
          </Form.Item>

          <Form.Item
            label="门店名称"
            name="dealer_name"
            rules={[{ required: true, message: '请输入门店名称' }]}
          >
            <Input placeholder="请输入门店名称" />
          </Form.Item>
        </div>

        <div className="grid grid-cols-2 gap-4">
          <Form.Item
            label="订单金额"
            name="amount"
            rules={[{ required: true, message: '请输入订单金额' }]}
          >
            <InputNumber
              placeholder="请输入订单金额"
              min={0}
              precision={2}
              className="w-full"
              addonBefore="¥"
            />
          </Form.Item>

          <Form.Item
            label="支付方式"
            name="pay_type"
          >
            <Input placeholder="请输入支付方式（可选）" />
          </Form.Item>
        </div>

        <Form.Item
          label="启用Mock"
          name="enable_mock"
          valuePropName="checked"
        >
          <Switch checkedChildren="启用" unCheckedChildren="停用" />
        </Form.Item>

        <div className="mb-4">
          <div className="flex items-center justify-between mb-3">
            <span className="text-base font-medium">商品列表</span>
            <Button
              type="dashed"
              icon={<PlusOutlined />}
              onClick={addProduct}
            >
              添加商品
            </Button>
          </div>

          <div className="space-y-3">
            {products.map((product, index) => (
              <Card
                key={index}
                size="small"
                className="bg-gray-50"
                extra={
                  products.length > 1 && (
                    <Button
                      type="text"
                      danger
                      size="small"
                      icon={<DeleteOutlined />}
                      onClick={() => removeProduct(index)}
                    />
                  )
                }
              >
                <div className="grid grid-cols-12 gap-3 items-center">
                  <div className="col-span-5">
                    <Input
                      placeholder="商品ID"
                      value={product.product_id}
                      onChange={(e) => updateProduct(index, 'product_id', e.target.value)}
                    />
                  </div>
                  <div className="col-span-5">
                    <Input
                      placeholder="商品名称"
                      value={product.name}
                      onChange={(e) => updateProduct(index, 'name', e.target.value)}
                    />
                  </div>
                  <div className="col-span-2">
                    <InputNumber
                      placeholder="数量"
                      min={1}
                      value={product.quantity}
                      onChange={(value) => updateProduct(index, 'quantity', value || 1)}
                      className="w-full"
                    />
                  </div>
                </div>
              </Card>
            ))}
          </div>
        </div>
      </Form>
    </Modal>
  );
};

export default AddConfigModal;
