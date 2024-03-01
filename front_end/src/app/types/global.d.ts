declare global {
  type TImage = {
    id?: string;
    img_id: string;
    img_name: string;
    mimetype: string;
    data?: Blob;
  };
  type TSpecks = {
    cpu?: string;
    featured?: string;
    io_ports?: string;
    more?: string[];
  };
  type TProduct = {
    id: string;
    product_id: string;
    product_model: string;
    reviews: number;
    name: string;
    images?: string[];
    isStock: boolean;
    brand: string;
    price: number;
    discount_price?: number;
    description: string;
    specks: TSpecks;
  };
}
export {};
