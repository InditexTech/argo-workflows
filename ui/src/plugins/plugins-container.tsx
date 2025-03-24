import * as React from 'react';
import {Route, RouteComponentProps, Switch} from 'react-router';
<<<<<<< HEAD:ui/src/app/reports/components/report-container.tsx
import {Loading} from '../../shared/components/loading';
=======
>>>>>>> draft-3.6.5:ui/src/plugins/plugins-container.tsx

import {PluginList} from './plugin-list';

export const PluginsContainer = (props: RouteComponentProps<any>) => (
    <Switch>
<<<<<<< HEAD:ui/src/app/reports/components/report-container.tsx
        <Route exact={true} path={`${props.match.path}/:namespace?`} component={SuspenseReports} />
=======
        <Route exact={true} path={`${props.match.path}/:namespace?`} component={PluginList} />
>>>>>>> draft-3.6.5:ui/src/plugins/plugins-container.tsx
    </Switch>
);

// lazy load Reports as it is infrequently used and imports large Chart components (which can be split into a separate bundle)
const LazyReports = React.lazy(() => import(/* webpackChunkName: "reports" */ './reports'));

function SuspenseReports(props: RouteComponentProps<any>) {
    return (
        <React.Suspense fallback={<Loading />}>
            <LazyReports {...props} />
        </React.Suspense>
    );
}
