
import { useState, useEffect } from 'react';
import { Layout, Button, Row, Col, Typography, Space, message, Spin, Empty, Modal, Input, Form } from 'antd';
import { PlusOutlined, ReloadOutlined, LockOutlined, LogoutOutlined } from '@ant-design/icons';
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
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [authModalVisible, setAuthModalVisible] = useState(false);
  const [authForm] = Form.useForm();

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

  // 检查认证状态
  useEffect(() => {
    const savedAuth = localStorage.getItem('order-mock-auth');
    if (savedAuth === 'zqf0517') {
      setIsAuthenticated(true);
      loadConfigs();
    } else {
      setAuthModalVisible(true);
    }
  }, []);

  // 处理密码验证
  const handleAuth = async () => {
    try {
      const values = await authForm.validateFields();
      const password = values.password;
      
      if (password === 'zqf0517') {
        localStorage.setItem('order-mock-auth', password);
        setIsAuthenticated(true);
        setAuthModalVisible(false);
        message.success('验证成功');
        loadConfigs();
      } else {
        message.error('密码错误，请重试');
        authForm.resetFields();
      }
    } catch (error) {
      console.error('验证失败:', error);
    }
  };

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

  // 退出登录
  const handleLogout = () => {
    localStorage.removeItem('order-mock-auth');
    setIsAuthenticated(false);
    setAuthModalVisible(true);
    setConfigs([]);
    message.info('已退出登录');
  };

  // 如果未认证，显示认证弹窗
  if (!isAuthenticated) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <Modal
          title={
            <div className="flex items-center gap-2">
              <LockOutlined className="text-blue-500" />
              <span>系统验证</span>
            </div>
          }
          open={authModalVisible}
          onOk={handleAuth}
          onCancel={() => {}}
          closable={false}
          maskClosable={false}
          okText="验证"
          cancelText="取消"
          width={400}
        >
          <div className="py-4">
            <p className="text-gray-600 mb-4">请输入访问密码以继续使用系统</p>
            <Form form={authForm} layout="vertical">
              <Form.Item
                name="password"
                label="访问密码"
                rules={[{ required: true, message: '请输入访问密码' }]}
              >
                <Input.Password
                  placeholder="请输入访问密码"
                  size="large"
                  onPressEnter={handleAuth}
                  autoFocus
                />
              </Form.Item>
            </Form>
          </div>
        </Modal>
        
        <div className="text-center">
          <LockOutlined className="text-6xl text-gray-300 mb-4" />
          <h2 className="text-xl text-gray-500">系统已锁定</h2>
          <p className="text-gray-400">请输入密码以访问订单Mock配置管理系统</p>
        </div>
      </div>
    );
  }

  return (
    <Layout className="min-h-screen bg-gray-50">
      <Header className="bg-white shadow-sm border-b">
        <div className="flex items-center justify-between h-full max-w-7xl mx-auto px-4">
          <div className="flex items-center space-x-3">
            <Title style={{ color: 'white' }} level={3} className="m-0 text-white">
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
            <Button
              icon={<LogoutOutlined />}
              onClick={handleLogout}
              danger
            >
              退出
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
                  管理门店订单Mock配置, 可以启用或停用特定门店的Mock功能
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
