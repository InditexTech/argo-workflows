import * as React from 'react';
import {ReactNode} from 'react';

import {Nudge} from './nudge';

export const SecurityNudge = (props: {children: ReactNode}) => (
    <Nudge key='security-nudge'>
        <i className='fa fa-lock-open status-icon--failed' /> {props.children}{' '}
<<<<<<< HEAD:ui/src/app/shared/components/security-nudge.tsx
        <a href='https://argo-workflows.readthedocs.io/en/release-3.5/workflow-pod-security-context/'>Learn more</a>
=======
        <a href='https://argo-workflows.readthedocs.io/en/latest/workflow-pod-security-context/'>Learn more</a>
>>>>>>> draft-3.6.5:ui/src/shared/components/security-nudge.tsx
    </Nudge>
);
