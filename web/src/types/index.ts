// 产品类型
export interface Product {
  product_id: string;
  name: string;
  quantity: number;
}

// 订单配置类型
export interface OrderConfig {
  dealer: string;
  dealer_name: string;
  enable_mock: boolean;
  products: Product[];
  amount: number;
  pay_type: string;
}

// API 响应类型
export interface ApiResponse<T> {
  code: number;
  message: string;
  data?: T;
}

// 删除配置参数
export interface DeleteConfigParams {
  dealer_id: string;
}
