// This file is autogenerated. DO NOT EDIT.
//
// To regenerate, run "make gen-cue" from repository root.
//
// Derived from the Thema lineage at pkg/coremodel/dashboard


// This model is a WIP and not yet canonical. Consequently, its members are
// not exported to exclude it from grafana-schema's public API surface.

interface AnnotationQuery {
  builtIn: number;
  datasource: {};
  enable: boolean;
  hide?: boolean;
  iconColor?: string;
  name?: string;
  rawQuery?: string;
  showIn: number;
  target?: {};
  type: string;
}

const defaultAnnotationQuery: Partial<AnnotationQuery> = {
  builtIn: 0,
  enable: true,
  hide: false,
  showIn: 0,
  type: 'dashboard',
};

interface VariableModel {
  label?: string;
  name: string;
  type: VariableType;
}

interface DashboardLink {
  asDropdown: boolean;
  icon?: string;
  includeVars: boolean;
  keepTime: boolean;
  tags: string[];
  targetBlank: boolean;
  title: string;
  tooltip?: string;
  type: DashboardLinkType;
  url?: string;
}

const defaultDashboardLink: Partial<DashboardLink> = {
  asDropdown: false,
  includeVars: false,
  keepTime: false,
  tags: [],
  targetBlank: false,
};

type DashboardLinkType = 'link' | 'dashboards';

type VariableType = 'query' | 'adhoc' | 'constant' | 'datasource' | 'interval' | 'textbox' | 'custom' | 'system';

enum FieldColorModeId {
  ContinuousGrYlRd = 'continuous-GrYlRd',
  Fixed = 'fixed',
  PaletteClassic = 'palette-classic',
  PaletteSaturated = 'palette-saturated',
  Thresholds = 'thresholds',
}

type FieldColorSeriesByMode = 'min' | 'max' | 'last';

interface FieldColor {
  fixedColor?: string;
  mode: FieldColorModeId | string;
  seriesBy?: FieldColorSeriesByMode;
}

interface Threshold {
  color: string;
  state?: string;
  value?: number;
}

enum ThresholdsMode {
  Absolute = 'absolute',
  Percentage = 'percentage',
}

interface ThresholdsConfig {
  mode: ThresholdsMode;
  steps: Threshold[];
}

const defaultThresholdsConfig: Partial<ThresholdsConfig> = {
  steps: [],
};

interface Transformation {
  id: string;
  options: {};
}

enum DashboardCursorSync {
  Crosshair = 1,
  Off = 0,
  Tooltip = 2,
}

const defaultDashboardCursorSync: DashboardCursorSync = DashboardCursorSync.Off;

interface Panel {
  datasource?: {};
  description?: string;
  fieldConfig: {
    defaults: {};
    overrides: {
      matcher: {
        id: string;
      };
      properties: {
        id: string;
      }[];
    }[];
  };
  gridPos?: {
    h: number;
    w: number;
    x: number;
    y: number;
  };
  id?: number;
  interval?: string;
  links?: DashboardLink[];
  maxDataPoints?: number;
  options: {};
  pluginVersion?: string;
  repeat?: string;
  repeatDirection: 'h' | 'v';
  tags?: string[];
  targets?: {}[];
  thresholds?: any[];
  timeFrom?: string;
  timeRegions?: any[];
  timeShift?: string;
  title?: string;
  transformations: Transformation[];
  transparent: boolean;
  type: string;
}

const defaultPanel: Partial<Panel> = {
  links: [],
  repeatDirection: 'h',
  tags: [],
  targets: [],
  thresholds: [],
  timeRegions: [],
  transformations: [],
  transparent: false,
};

interface Dashboard {
  annotations?: {
    list: AnnotationQuery[];
  };
  description?: string;
  editable: boolean;
  fiscalYearStartMonth?: number;
  gnetId?: string;
  graphTooltip: DashboardCursorSync;
  id?: number;
  links?: DashboardLink[];
  liveNow?: boolean;
  panels?: Panel | {
      type: 'graph';
    } | {
      type: 'heatmap';
    } | {
      type: 'row';
      collapsed: boolean;
      id: number;
      panels: Panel | {
          type: 'graph';
        } | {
          type: 'heatmap';
        }[];
    }[];
  refresh?: string | false;
  schemaVersion: number;
  style: 'light' | 'dark';
  tags?: string[];
  templating?: {
    list: VariableModel[];
  };
  time?: {
    from: string;
    to: string;
  };
  timepicker?: {
    collapse: boolean;
    enable: boolean;
    hidden: boolean;
    refresh_intervals: string[];
  };
  timezone?: 'browser' | 'utc' | '';
  title?: string;
  uid?: string;
  version?: number;
  weekStart?: string;
}

const defaultDashboard: Partial<Dashboard> = {
  editable: true,
  graphTooltip: DashboardCursorSync.Off,
  links: [],
  panels: [],
  schemaVersion: 36,
  style: 'dark',
  tags: [],
  timezone: 'browser',
};
