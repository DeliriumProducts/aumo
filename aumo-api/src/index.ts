export * from './lib/aumo';
export * from './lib/auth';
export * from './lib/config';
export * from './lib/order';
export * from './lib/product';
export * from './lib/receipt';
export * from './lib/user';

import auth from './lib/auth';
import config from './lib/config';
import order from './lib/order';
import product from './lib/product';
import receipt from './lib/receipt';
import user from './lib/user';

export default {
  auth,
  config,
  order,
  product,
  receipt,
  user
};
