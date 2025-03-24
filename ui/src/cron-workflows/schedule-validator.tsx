import x from 'cronstrue';
import * as React from 'react';
<<<<<<< HEAD:ui/src/app/cron-workflows/components/schedule-validator.tsx

import {SuccessIcon, WarningIcon} from '../../shared/components/fa-icons';

=======

import {SuccessIcon, WarningIcon} from '../shared/components/fa-icons';

>>>>>>> draft-3.6.5:ui/src/cron-workflows/schedule-validator.tsx
export function ScheduleValidator({schedule}: {schedule: string}) {
    try {
        if (schedule.split(' ').length >= 6) {
            throw new Error('cron schedules must consist of 5 values only');
        }
        return (
            <span>
                <SuccessIcon /> {x.toString(schedule)}
            </span>
        );
    } catch (e) {
        return (
            <span>
                <WarningIcon /> Schedule maybe invalid: {e.toString()}
            </span>
        );
    }
}
