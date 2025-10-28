import type { OrderConfig, ApiResponse, DeleteConfigParams } from '../types';

const API_BASE_URL = '/api';

// 获取所有配置
export const getMockConfigs = async (): Promise<OrderConfig[]> => {
  try {
    const response = await fetch(`${API_BASE_URL}/mock_config`);
    const data: ApiResponse<OrderConfig[]> = await response.json();
    
    if (data.code === 0) {
      return data.data || [];
    } else {
      throw new Error(data.message || '获取配置失败');
    }
  } catch (error) {
    console.error('获取配置失败:', error);
    throw error;
  }
};

// 添加门店配置
export const addDealerConfig = async (config: OrderConfig): Promise<void> => {
  try {
    const response = await fetch(`${API_BASE_URL}/add_dealer`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(config),
    });
    
    const data: ApiResponse<null> = await response.json();
    
    if (data.code !== 0) {
      throw new Error(data.message || '添加配置失败');
    }
  } catch (error) {
    console.error('添加配置失败:', error);
    throw error;
  }
};

// 删除门店配置
export const deleteDealerConfig = async (dealerId: string): Promise<void> => {
  try {
    const params: DeleteConfigParams = { dealer_id: dealerId };
    const response = await fetch(`${API_BASE_URL}/delete_dealer_config`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(params),
    });
    
    const data: ApiResponse<null> = await response.json();
    
    if (data.code !== 0) {
      throw new Error(data.message || '删除配置失败');
    }
  } catch (error) {
    console.error('删除配置失败:', error);
    throw error;
  }
};
