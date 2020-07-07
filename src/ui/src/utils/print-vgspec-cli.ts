/* eslint-disable no-console */
import { DARK_THEME } from 'common/mui-theme';
import { Data } from 'vega';

import {
  BAR_CHART_TYPE,
  ChartDisplay,
  convertWidgetDisplayToVegaSpec,
  TIMESERIES_CHART_TYPE,
} from 'containers/live/convert-to-vega-spec';
import { DISPLAY_TYPE_KEY } from 'containers/live/vis';

const timeseriesData = [
  { time_: '4/2/2020, 9:42:38 PM', service: 'px-sock-shop/catalogue', bytesPerSecond: 48259 },
  { time_: '4/2/2020, 9:42:38 PM', service: 'px-sock-shop/orders', bytesPerSecond: 234 },
  { time_: '4/2/2020, 9:42:39 PM', service: 'px-sock-shop/catalogue', bytesPerSecond: 52234 },
  { time_: '4/2/2020, 9:42:39 PM', service: 'px-sock-shop/orders', bytesPerSecond: 23423 },
  { time_: '4/2/2020, 9:42:40 PM', service: 'px-sock-shop/catalogue', bytesPerSecond: 18259 },
  { time_: '4/2/2020, 9:42:40 PM', service: 'px-sock-shop/orders', bytesPerSecond: 28259 },
  { time_: '4/2/2020, 9:42:41 PM', service: 'px-sock-shop/catalogue', bytesPerSecond: 38259 },
  { time_: '4/2/2020, 9:42:42 PM', service: 'px-sock-shop/orders', bytesPerSecond: 10259 },
  { time_: '4/2/2020, 9:42:43 PM', service: 'px-sock-shop/catalogue', bytesPerSecond: 58259 },
];

const barData = [
  {
    service: 'carts', endpoint: '/create', cluster: 'prod', numErrors: 14,
  },
  {
    service: 'carts', endpoint: '/create', cluster: 'staging', numErrors: 60,
  },
  {
    service: 'carts', endpoint: '/create', cluster: 'dev', numErrors: 3,
  },
  {
    service: 'carts', endpoint: '/create', cluster: 'prod', numErrors: 80,
  },
  {
    service: 'carts', endpoint: '/create', cluster: 'staging', numErrors: 38,
  },
  {
    service: 'carts', endpoint: '/update', cluster: 'dev', numErrors: 55,
  },
  {
    service: 'carts', endpoint: '/submit', cluster: 'prod', numErrors: 11,
  },
  {
    service: 'carts', endpoint: '/submit', cluster: 'staging', numErrors: 58,
  },
  {
    service: 'carts', endpoint: '/submit', cluster: 'dev', numErrors: 79,
  },
  {
    service: 'orders', endpoint: '/remove', cluster: 'prod', numErrors: 83,
  },
  {
    service: 'orders', endpoint: '/remove', cluster: 'staging', numErrors: 87,
  },
  {
    service: 'orders', endpoint: '/remove', cluster: 'dev', numErrors: 67,
  },
  {
    service: 'orders', endpoint: '/add', cluster: 'prod', numErrors: 97,
  },
  {
    service: 'orders', endpoint: '/add', cluster: 'staging', numErrors: 84,
  },
  {
    service: 'orders', endpoint: '/add', cluster: 'dev', numErrors: 90,
  },
  {
    service: 'orders', endpoint: '/add', cluster: 'prod', numErrors: 74,
  },
  {
    service: 'orders', endpoint: '/new', cluster: 'staging', numErrors: 64,
  },
  {
    service: 'orders', endpoint: '/new', cluster: 'dev', numErrors: 19,
  },
  {
    service: 'frontend', endpoint: '/orders', cluster: 'prod', numErrors: 57,
  },
  {
    service: 'frontend', endpoint: '/orders', cluster: 'staging', numErrors: 35,
  },
  {
    service: 'frontend', endpoint: '/orders', cluster: 'dev', numErrors: 49,
  },
  {
    service: 'frontend', endpoint: '/redirect', cluster: 'prod', numErrors: 91,
  },
  {
    service: 'frontend', endpoint: '/signup', cluster: 'staging', numErrors: 38,
  },
  {
    service: 'frontend', endpoint: '/signup', cluster: 'dev', numErrors: 91,
  },
  {
    service: 'frontend', endpoint: '/signup', cluster: 'prod', numErrors: 99,
  },
  {
    service: 'frontend', endpoint: '/signup', cluster: 'staging', numErrors: 80,
  },
  {
    service: 'frontend', endpoint: '/signup', cluster: 'dev', numErrors: 37,
  },
];

function printSpec(display: ChartDisplay) {
  const sourceName = 'mysource';
  const { spec } = convertWidgetDisplayToVegaSpec(display, sourceName, DARK_THEME);
  let data: Array<{}>;

  if (display[DISPLAY_TYPE_KEY] === BAR_CHART_TYPE) {
    data = barData;
  } else if (display[DISPLAY_TYPE_KEY] === TIMESERIES_CHART_TYPE) {
    data = timeseriesData;
  } else {
    console.log(
      `This tool only supports bar and timeseries charts, not ${display[DISPLAY_TYPE_KEY]}`);
    return;
  }

  for (const datum of (spec.data as Data[])) {
    if (datum.name === sourceName) {
      (datum as any).values = data;
    }
  }

  // Remove everything that uses custom extensions to vega so that the spec works out of the box in
  // the online vega editor.
  (spec as any).signals = (spec as any).signals.filter((signal) => !signal.name.includes('ts_domain_value'));
  if ((spec as any).scales) {
    for (const scale of (spec as any).scales) {
      if (scale.name === 'x' && scale.domainRaw) {
        delete scale.domainRaw;
      }
    }
  }
  if ((spec as any).axes) {
    for (const axis of (spec as any).axes) {
      if (axis.scale === 'x' && axis.encode && axis.encode.labels && axis.encode.labels.update
          && axis.encode.labels.update.text && axis.encode.labels.update.text.signal
          && axis.encode.labels.update.text.signal.includes('pxTimeFormat')) {
        // TODO(james): figure out how to get the x axis labels to show something reasonable without
        // pxTimeFormat.
        axis.encode.labels.update.text.signal = '';
      }
    }
  }

  console.log(JSON.stringify(spec));
}

// Example input, replace with what you want to print the spec for.
const input = {
  '@type': TIMESERIES_CHART_TYPE,
  timeseries: [{
    value: 'bytesPerSecond',
    mode: 'MODE_AREA',
    series: 'service',
    stackBySeries: true,
  }],
};

printSpec(input);

/* eslint-enable no-console */