import { useCallback, useEffect, useRef, useState } from "react";

/**
 * Voice input via the browser's SpeechRecognition API.
 *
 * Deliberately browser-native rather than a Whisper pipeline: it needs no
 * model download, no backend round trip and no extra dependency, so dictation
 * works the moment the page loads. The trade-off is that it is Chrome/Edge/
 * Safari only and sends audio to the browser vendor's service, which is why
 * `supported` is exposed and the UI hides the control rather than showing a
 * button that silently does nothing.
 *
 * A local Whisper path can later implement this same hook signature for users
 * who need dictation to stay on-device.
 */

interface SpeechRecognitionAlternativeLike {
  transcript: string;
}
interface SpeechRecognitionResultLike {
  0: SpeechRecognitionAlternativeLike;
  isFinal: boolean;
  length: number;
}
interface SpeechRecognitionEventLike {
  resultIndex: number;
  results: {
    length: number;
    [index: number]: SpeechRecognitionResultLike;
  };
}
interface SpeechRecognitionLike {
  lang: string;
  continuous: boolean;
  interimResults: boolean;
  start(): void;
  stop(): void;
  abort(): void;
  onresult: ((e: SpeechRecognitionEventLike) => void) | null;
  onerror: ((e: { error: string }) => void) | null;
  onend: (() => void) | null;
}

type Ctor = new () => SpeechRecognitionLike;

function getConstructor(): Ctor | undefined {
  const w = window as unknown as {
    SpeechRecognition?: Ctor;
    webkitSpeechRecognition?: Ctor;
  };
  return w.SpeechRecognition ?? w.webkitSpeechRecognition;
}

export interface VoiceInput {
  supported: boolean;
  listening: boolean;
  /** Text recognised so far this session, including the interim guess. */
  transcript: string;
  error?: string;
  start(lang?: string): void;
  stop(): void;
  reset(): void;
}

export function useVoiceInput(onFinal?: (text: string) => void): VoiceInput {
  const [supported] = useState(() => !!getConstructor());
  const [listening, setListening] = useState(false);
  const [transcript, setTranscript] = useState("");
  const [error, setError] = useState<string>();

  const recognitionRef = useRef<SpeechRecognitionLike | null>(null);
  const finalRef = useRef("");
  // Keep the latest callback without restarting recognition on every render.
  const onFinalRef = useRef(onFinal);
  useEffect(() => {
    onFinalRef.current = onFinal;
  }, [onFinal]);

  const stop = useCallback(() => {
    recognitionRef.current?.stop();
    setListening(false);
  }, []);

  const reset = useCallback(() => {
    finalRef.current = "";
    setTranscript("");
    setError(undefined);
  }, []);

  const start = useCallback(
    (lang = navigator.language || "en-US") => {
      const Ctor = getConstructor();
      if (!Ctor) {
        setError("Voice input is not supported in this browser.");
        return;
      }

      // Restarting cleanly avoids the "already started" exception.
      recognitionRef.current?.abort();

      const rec = new Ctor();
      rec.lang = lang;
      rec.continuous = true;
      rec.interimResults = true;

      rec.onresult = (e) => {
        let interim = "";
        for (let i = e.resultIndex; i < e.results.length; i++) {
          const result = e.results[i];
          const text = result[0].transcript;
          if (result.isFinal) finalRef.current += text;
          else interim += text;
        }
        setTranscript((finalRef.current + interim).trimStart());
      };

      rec.onerror = (e) => {
        setError(
          e.error === "not-allowed"
            ? "Microphone permission was denied."
            : e.error === "no-speech"
              ? "Didn't catch that, try again."
              : `Voice input failed (${e.error}).`,
        );
        setListening(false);
      };

      rec.onend = () => {
        setListening(false);
        const final = finalRef.current.trim();
        if (final) onFinalRef.current?.(final);
      };

      recognitionRef.current = rec;
      finalRef.current = "";
      setTranscript("");
      setError(undefined);
      setListening(true);
      rec.start();
    },
    [],
  );

  useEffect(() => () => recognitionRef.current?.abort(), []);

  return { supported, listening, transcript, error, start, stop, reset };
}
