import type { Product } from './product'
import type { Metadata } from './metadata'

export interface ProductsListResponse {
  data?: Product[]
  metadata?: Metadata
}
