import * as React from 'react';
import {ReactNode} from 'react';

import {Nudge} from './nudge';

export const CostOptimisationNudge = (props: {name: string; children: ReactNode}) => (
    <Nudge key={'cost-optimization-nudge/' + props.name}>
<<<<<<< HEAD:ui/src/app/shared/components/cost-optimisation-nudge.tsx
        <i className='fa fa-money-bill-alt status-icon--pending' /> {props.children}{' '}
        <a href='https://argo-workflows.readthedocs.io/en/release-3.5/cost-optimisation/'>Learn more</a>
=======
        <i className='fa fa-money-bill-alt status-icon--pending' /> {props.children} <a href='https://argo-workflows.readthedocs.io/en/latest/cost-optimisation/'>Learn more</a>
>>>>>>> draft-3.6.5:ui/src/shared/components/cost-optimisation-nudge.tsx
    </Nudge>
);
