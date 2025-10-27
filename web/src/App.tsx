
import React, { useState, useEffect } from 'react';
import { Layout, Button, Row, Col, Typography, Space, message, Spin, Empty } from 'antd';
import { PlusOutlined, ReloadOutlined, SettingOutlined } from '@ant-design/icons';
import ConfigCard from './components/ConfigCard';
import AddConfigModal from './components/AddConfigModal';
import type { OrderConfig } from './types';
import { getMockConfigs, addDealerConfig, deleteDealerConfig } from './services/api';
import './assets/styles/index.css';

const { Header, Content } = Layout;
const { Title } = Typography;

function App() {
  const [configs, setConfigs] = useState<OrderConfig[]>([]);
  const [loading, setLoading] = useState(false);
  const [modalVisible, setModalVisible] = useState(false);
  const [editingConfig, setEditingConfig] = useState<OrderConfig | null>(null);

  // 加载配置数据
  const loadConfigs = async () => {
    try {
      setLoading(true);
      const data = await getMockConfigs();
      setConfigs(data);
    } catch (error) {
      message.error('加载配置失败');
      console.error('加载配置失败:', error);
    } finally {
      setLoading(false);
    }
  };

  // 初始化加载
  useEffect(() => {
    loadConfigs();
  }, []);

  // 处理添加配置
  const handleAddConfig = async (config: OrderConfig) => {
    await addDealerConfig(config);
    await loadConfigs();
  };

  // 处理删除配置
  const handleDeleteConfig = async (dealerId: string) => {
    try {
      await deleteDealerConfig(dealerId);
      message.success('删除配置成功');
      await loadConfigs();
    } catch (error) {
      message.error('删除配置失败');
      console.error('删除配置失败:', error);
    }
  };

  // 处理Mock开关切换
  const handleToggleMock = async (dealerId: string, enabled: boolean) => {
    try {
      // 这里暂时只是前端状态更新，后端接口还未实现
      const updatedConfigs = configs.map(config =>
        config.dealer === dealerId
          ? { ...config, enable_mock: enabled }
          : config
      );
      setConfigs(updatedConfigs);
      message.success(`已${enabled ? '启用' : '停用'}Mock功能`);
      
      // TODO: 调用后端API更新Mock状态
      console.log(`Toggle mock for dealer ${dealerId}: ${enabled}`);
    } catch (error) {
      message.error('切换Mock状态失败');
      console.error('切换Mock状态失败:', error);
    }
  };

  // 处理编辑配置
  const handleEditConfig = (config: OrderConfig) => {
    setEditingConfig(config);
    setModalVisible(true);
  };

  // 关闭弹窗
  const handleModalCancel = () => {
    setModalVisible(false);
    setEditingConfig(null);
  };

  return (
    <Layout className="min-h-screen bg-gray-50">
      <Header className="bg-white shadow-sm border-b">
        <div className="flex items-center justify-between h-full max-w-7xl mx-auto px-4">
          <div className="flex items-center space-x-3">
            <SettingOutlined className="text-2xl text-blue-500" />
            <Title level={3} className="m-0 text-gray-800">
              订单Mock配置管理
            </Title>
          </div>
          <Space>
            <Button
              icon={<ReloadOutlined />}
              onClick={loadConfigs}
              loading={loading}
            >
              刷新
            </Button>
            <Button
              type="primary"
              icon={<PlusOutlined />}
              onClick={() => setModalVisible(true)}
            >
              添加配置
            </Button>
          </Space>
        </div>
      </Header>

      <Content className="p-6">
        <div className="max-w-7xl mx-auto">
          <div className="mb-6">
            <div className="flex items-center justify-between">
              <div>
                <Title level={4} className="m-0">
                  配置列表
                </Title>
                <p className="text-gray-500 mt-1">
                  管理门店订单Mock配置，可以启用或停用特定门店的Mock功能
                </p>
              </div>
              <div className="text-sm text-gray-500">
                共 {configs.length} 个配置
              </div>
            </div>
          </div>

          <Spin spinning={loading}>
            {configs.length === 0 ? (
              <div className="text-center py-12">
                <Empty
                  description="暂无配置数据"
                  image={Empty.PRESENTED_IMAGE_SIMPLE}
                >
                  <Button
                    type="primary"
                    icon={<PlusOutlined />}
                    onClick={() => setModalVisible(true)}
                  >
                    添加第一个配置
                  </Button>
                </Empty>
              </div>
            ) : (
              <Row gutter={[16, 16]}>
                {configs.map((config) => (
                  <Col xs={24} sm={12} lg={8} xl={6} key={config.dealer}>
                    <ConfigCard
                      config={config}
                      onToggleMock={handleToggleMock}
                      onDelete={handleDeleteConfig}
                      onEdit={handleEditConfig}
                    />
                  </Col>
                ))}
              </Row>
            )}
          </Spin>
        </div>
      </Content>

      <AddConfigModal
        visible={modalVisible}
        onCancel={handleModalCancel}
        onSubmit={handleAddConfig}
        editingConfig={editingConfig}
      />
    </Layout>
  );
}

export default App;
