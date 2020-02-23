export interface AumoOptions {
  Backend: string;
}

export let options: AumoOptions = {
  Backend: 'https://aumo-api.deliriumproducts.me/api/v1'
};

export function config(opts: AumoOptions) {
  options = opts;
}

export default { config, options };
