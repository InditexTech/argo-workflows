import * as React from 'react';

import {denominator, formatDuration} from './duration';

interface Props {
    resourcesDuration: {[resource: string]: number};
}

export function ResourcesDuration(props: Props) {
    return (
        <>
            {props.resourcesDuration &&
                Object.entries(props.resourcesDuration)
                    .map(([resource, duration]) => formatDuration(duration, 1) + '*(' + denominator(resource) + ' ' + resource + ')')
                    .join(',')}{' '}
<<<<<<< HEAD:ui/src/app/shared/resources-duration.tsx
            <a href='https://argo-workflows.readthedocs.io/en/release-3.5/resource-duration/'>
=======
            <a href='https://argo-workflows.readthedocs.io/en/latest/resource-duration/'>
>>>>>>> draft-3.6.5:ui/src/shared/resources-duration.tsx
                <i className='fa fa-info-circle' />
            </a>
        </>
    );
}
