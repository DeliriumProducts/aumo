export interface AumoOptions {
  Backend: string;
}

export let options: AumoOptions = {
  Backend: 'https://aumo-api.deliprods.tech/api/v1'
};

export function config(opts: AumoOptions) {
  options = opts;
}

export default { config, options };
