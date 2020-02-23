export * from './lib/aumo';
export * from './lib/auth';
export * from './lib/config';
export * from './lib/order';
export * from './lib/receipt';
export * from './lib/user';

import auth from './lib/auth';
import config from './lib/config';
import order from './lib/order';
import receipt from './lib/receipt';
import shop from './lib/shop';
import user from './lib/user';

export default {
  auth,
  config,
  order,
  receipt,
  user,
  shop
};
