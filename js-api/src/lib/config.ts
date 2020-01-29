export interface AumoOptions {
  Backend: string;
}

export let options: AumoOptions = {
  Backend: 'https://aumo-api.deliriumproducts.me'
};

export function config(opts: AumoOptions) {
  options = opts;
}
