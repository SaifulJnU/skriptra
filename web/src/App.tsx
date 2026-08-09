import { Navigate, Route, Routes } from "react-router-dom";
import AppShell from "@/components/AppShell";
import Dashboard from "@/pages/Dashboard";
import CourseOverview from "@/pages/CourseOverview";
import Exams from "@/pages/Exams";
import Questions from "@/pages/Questions";
import QuestionViewer from "@/pages/QuestionViewer";
import Ask from "@/pages/Ask";
import Analytics from "@/pages/Analytics";
import NotFound from "@/pages/NotFound";

export default function App() {
  return (
    <Routes>
      <Route element={<AppShell />}>
        <Route path="/" element={<Dashboard />} />
        <Route path="/courses/:courseId" element={<CourseOverview />} />
        <Route path="/courses/:courseId/exams" element={<Exams />} />
        <Route path="/courses/:courseId/questions" element={<Questions />} />
        <Route path="/courses/:courseId/ask" element={<Ask />} />
        <Route path="/courses/:courseId/analytics" element={<Analytics />} />
        <Route path="/questions/:questionId" element={<QuestionViewer />} />
        <Route path="/index.html" element={<Navigate to="/" replace />} />
        <Route path="*" element={<NotFound />} />
      </Route>
    </Routes>
  );
}
