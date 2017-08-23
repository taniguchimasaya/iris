// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package runtime

import "testing"

func TestSignalCondition(t *testing.T) {
	defspecial(Defun)
	defun(Cerror)
	defun(ContinueCondition)
	defspecial(WithHandler)
	defspecial(Function)
	defspecial(Quote)
	tests := []test{
		{
			exp: ` 
				(defun continue-condition-handler (condition)
					(continue-condition condition 999))
				`,
			want:    `'continue-condition-handler`,
			wantErr: false,
		},
		{
			exp: ` 
				 (with-handler #'continue-condition-handler
					(cerror "cont" "err")) 
				`,
			want:    `999`,
			wantErr: false,
		},
		{
			exp: ` 
				 (with-handler #'continue-condition-handler
					(error "not cont")) 
				`,
			want:    ``,
			wantErr: true,
		},
	}
	execTests(t, SignalCondition, tests)
}
