import * as React from 'react';
import {useEffect, useRef, useState} from 'react';
import MonacoEditor from 'react-monaco-editor';

<<<<<<< HEAD:ui/src/app/shared/components/object-editor/object-editor.tsx
import {uiUrl} from '../../base';
import {ScopedLocalStorage} from '../../scoped-local-storage';
import {Button} from '../button';
import {parse, stringify} from '../object-parser';
import {PhaseIcon} from '../phase-icon';
import {SuspenseMonacoEditor} from '../suspense-monaco-editor';
=======
import {uiUrl} from '../base';
import {useEditableObject} from '../use-editable-object';
import {Button} from './button';
import type {Lang} from './object-parser';
import {PhaseIcon} from './phase-icon';
import {SuspenseMonacoEditor} from './suspense-monaco-editor';
>>>>>>> draft-3.6.5:ui/src/shared/components/object-editor.tsx

interface Props<T> {
    type?: string;
    value: T;
    lang: Lang;
    text: string;
    onLangChange: (lang: Lang) => void;
    onChange?: (value: string) => void;
}

export function ObjectEditor<T>({type, value, text, lang, onChange, onLangChange}: Props<T>) {
    const [error, setError] = useState<Error>();
<<<<<<< HEAD:ui/src/app/shared/components/object-editor/object-editor.tsx
    const [lang, setLang] = useState<string>(storage.getItem('lang', defaultLang));
    const [text, setText] = useState<string>(stringify(value, lang));
=======
>>>>>>> draft-3.6.5:ui/src/shared/components/object-editor.tsx
    const editor = useRef<MonacoEditor>(null);

    useEffect(() => {
<<<<<<< HEAD:ui/src/app/shared/components/object-editor/object-editor.tsx
        if (!editor.current) {
            return;
        }

        // we ONLY want to change the text, if the normalized version has changed, this prevents white-space changes
        // from resulting in a significant change
        const editorText = stringify(parse(editor.current.editor.getValue()), lang);
        const editorLang = editor.current.editor.getValue().startsWith('{') ? 'json' : 'yaml';
        if (text !== editorText || lang !== editorLang) {
            editor.current.editor.setValue(stringify(parse(text), lang));
        }
    }, [editor, text, lang]);
=======
        if (!editor.current || text === editor.current.editor.getValue()) {
            return;
        }
        editor.current.editor.setValue(text);
    }, [editor, text]);
>>>>>>> draft-3.6.5:ui/src/shared/components/object-editor.tsx

    useEffect(() => {
        if (!type || lang !== 'json') {
            return;
        }

        (async () => {
            const uri = uiUrl('assets/jsonschema/schema.json');
            try {
                const res = await fetch(uri);
                const swagger = await res.json();
                // lazy load this, otherwise all of monaco-editor gets imported into the main bundle
                const languages = (await import(/* webpackChunkName: "monaco-editor" */ 'monaco-editor/esm/vs/editor/editor.api')).languages;
                // adds auto-completion to JSON only
                languages.json.jsonDefaults.setDiagnosticsOptions({
                    validate: true,
                    schemas: [
                        {
                            uri,
                            fileMatch: ['*'],
                            schema: {
                                $id: 'http://workflows.argoproj.io/' + type + '.json',
                                $ref: '#/definitions/' + type,
                                $schema: 'http://json-schema.org/draft-07/schema',
                                definitions: swagger.definitions
                            }
                        }
                    ]
                });
            } catch (err) {
                setError(err);
            }
        })();
    }, [lang, type]);

    // this calculation is rough, it is probably hard to work for for every case, essentially it is:
    // some pixels above and below for buttons, plus a bit of a buffer/padding
    const height = Math.max(600, window.innerHeight * 0.9 - 250);

    return (
        <>
            <div style={{paddingBottom: '1em'}}>
                <Button outline={true} onClick={() => onLangChange(lang === 'yaml' ? 'json' : 'yaml')}>
                    <span style={{fontWeight: lang === 'json' ? 'bold' : 'normal'}}>JSON</span>/<span style={{fontWeight: lang === 'yaml' ? 'bold' : 'normal'}}>YAML</span>
                </Button>

                {Object.keys(value).map(x => (
                    <Button
                        key={x}
                        icon='caret-right'
                        outline={true}
                        onClick={() => {
                            // Attempt to move the correct section of the document. Ideally, we'd have the line at the top of the
                            // editor, but Monaco editor does not have method for this (e.g. `revealLineAtTop`).

                            // find the line for the section in either YAML or JSON
                            const index = text.split('\n').findIndex(y => (lang === 'yaml' ? y.startsWith(x + ':') : y.includes('"' + x + '":')));

                            if (index >= 0) {
                                const lineNumber = index + 1;
                                editor.current.editor.revealLineInCenter(lineNumber);
                                editor.current.editor.setPosition({lineNumber, column: 0});
                                editor.current.editor.focus();
                            }
                        }}>
                        {x}
                    </Button>
                ))}
            </div>
            <div>
                <SuspenseMonacoEditor
                    ref={editor}
                    key='editor'
                    defaultValue={text}
                    language={lang}
                    height={height + 'px'}
                    options={{
                        readOnly: onChange === null,
                        minimap: {enabled: false},
                        guides: {
                            indentation: false
                        },
                        scrollBeyondLastLine: true
                    }}
                    onChange={v => {
                        if (onChange) {
                            try {
                                onChange(v);
                                setError(null);
                            } catch (e) {
                                setError(e);
                            }
                        }
                    }}
                />
            </div>
            {error && (
                <div style={{paddingTop: '1em'}}>
                    <PhaseIcon value='Error' /> {error.message}
                </div>
            )}
            {onChange && (
                <div>
                    <i className='fa fa-info-circle' />{' '}
                    {lang === 'json' ? <>Full auto-completion enabled.</> : <>Basic completion for YAML. Switch to JSON for full auto-completion.</>}{' '}
<<<<<<< HEAD:ui/src/app/shared/components/object-editor/object-editor.tsx
                    <a href='https://argo-workflows.readthedocs.io/en/release-3.5/ide-setup/'>Learn how to get auto-completion in your IDE.</a>
=======
                    <a href='https://argo-workflows.readthedocs.io/en/latest/ide-setup/'>Learn how to get auto-completion in your IDE.</a>
>>>>>>> draft-3.6.5:ui/src/shared/components/object-editor.tsx
                </div>
            )}
        </>
    );
}

/** Wrapper for ObjectEditor that automatically handles serializing/deserializing the object using useEditableObject() */
export function SerializingObjectEditor<T>({type, value}: {type?: string; value: T}) {
    const {object, setObject, serialization, lang, setLang} = useEditableObject<T>(value);
    return <ObjectEditor type={type} value={object} text={serialization} lang={lang} onLangChange={setLang} onChange={setObject} />;
}
