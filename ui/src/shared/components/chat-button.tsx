import * as React from 'react';
import {useEffect, useState} from 'react';

import {Link} from '../models';
import {services} from '../services';

export function ChatButton() {
    const [link, setLink] = useState<Link>();

    useEffect(() => {
        services.info
            .getInfo()
            .then(info => info.links)
<<<<<<< HEAD:ui/src/app/shared/components/chat-button.tsx
            .then(links => (links || []).concat({name: 'Get help', scope: 'chat', url: 'https://argo-workflows.readthedocs.io/en/release-3.5/'}).filter(x => x.scope === 'chat'))
=======
            .then(links => (links || []).concat({name: 'Get help', scope: 'chat', url: 'https://argo-workflows.readthedocs.io/en/latest/'}).filter(x => x.scope === 'chat'))
>>>>>>> draft-3.6.5:ui/src/shared/components/chat-button.tsx
            .then(links => {
                setLink(links[0]);
            });
    }, []);

    if (!link) {
        return null;
    }

    return (
        <div style={{position: 'fixed', right: 10, bottom: 10}}>
            <a href={link.url} className='argo-button argo-button--special'>
                <i className='fas fa-comment-alt' /> {link.name}
            </a>
        </div>
    );
}
